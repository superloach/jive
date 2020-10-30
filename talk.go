package jive

import (
	"errors"
	"log"
	"time"

	fastws "github.com/fasthttp/websocket"
	ws "github.com/gofiber/websocket/v2"
)

func (j *Jive) Talk(c *ws.Conn) {
	user := c.Params("user")
	name := c.Params("room")

	log.Println("[talk] begin", user, name)

	room := j.Room(name)

	conn, err := room.Open(user)
	if err != nil {
		log.Println("[talk] open err:", err)
		_ = c.WriteJSON(ErrEvent(user, err))
		return
	}

	log.Println("[talk] opened", user)

	go room.Do(HeyEvent(user))
	log.Println("[talk] did hey", user)

	go func() {
		for {
			evt := (*Event)(nil)

			err := c.ReadJSON(&evt)
			if err != nil {
				closeErr := (*fastws.CloseError)(nil)

				if errors.As(err, &closeErr) {
					log.Println("[talk] close", user)
					room.Close(user)
					return
				}

				log.Println("[talk] read err:", err)
				conn <- ErrEvent(user, err)
				continue
			}

			evt.When = time.Now()
			evt.User = user

			room.Do(evt)
		}
	}()

	for e := range conn {
		log.Println("[talk] conn got", e)

		err := c.WriteJSON(e)
		if err != nil {
			log.Println("[talk] write err:", err)
			continue
		}
	}

	room.Close(user)
}
