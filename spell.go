package jive

import "log"

type Spell func(r *Room, e *Event)

func TestSpell(r *Room, e *Event) {
	log.Println("[spell] test :3")
	r.Do(MsgEvent(e.User, "testing a spell :3"))
}
