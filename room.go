package jive

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

type Room struct {
	sync.Mutex

	Jive *Jive

	Name  string
	Conns map[string]chan *Event
}

func (j *Jive) Room(name string) *Room {
	j.Lock()
	defer j.Unlock()

	for _, room := range j.Rooms {
		if room.Name == name {
			log.Println("[room] got", name)
			return room
		}
	}

	log.Println("[room] new", name)

	room := &Room{
		Jive: j,

		Name:  name,
		Conns: map[string]chan *Event{},
	}
	j.Rooms = append(j.Rooms, room)
	return room
}

func (r *Room) Open(user string) (chan *Event, error) {
	r.Lock()
	defer r.Unlock()

	log.Println("[room] open", user)

	_, ok := r.Conns[user]
	if ok {
		return nil, fmt.Errorf("%s is already in this room")
	}

	conn := make(chan *Event)
	r.Conns[user] = conn

	return conn, nil
}

func (r *Room) Close(user string) {
	r.Lock()
	defer r.Unlock()

	log.Println("[room] close", r.Name, user)

	conn, ok := r.Conns[user]
	if !ok {
		log.Println("[room]", r.Name, user, "already closed")
		return
	}

	close(conn)
	delete(r.Conns, user)
}

func (r *Room) Do(e *Event) {
	if e.Type == "msg" && len(e.Msg) > 0 && e.Msg[0] == ':' {
		args := strings.Fields(e.Msg[1:])

		r.Do(SplEvent(
			e.User,
			args[0],
			args[1:]...,
		))
		return
	}

	log.Println("[room]", r.Name, "do", e)

	if e.Type == "spl" {
		spl := e.Spl[0]

		spell, ok := r.Jive.Spells[spl]
		if !ok {
			conn := r.Conns[e.User]
			conn <- ErrEvent(e.User, "unknown spell "+spl)
			return
		}

		spell(r, e)
		return
	}

	for user, conn := range r.Conns {
		log.Println("[room] conn", r.Name, user, "do", e)
		conn <- e
	}
}
