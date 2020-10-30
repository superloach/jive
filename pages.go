package jive

import (
	"log"

	f "github.com/gofiber/fiber/v2"
)

func (j *Jive) IndexPage(c *f.Ctx) error {
	log.Println("[pages] index")
	return c.Render("index", f.Map{
		"Rooms": j.Rooms,
	})
}

func (j *Jive) RoomPage(c *f.Ctx) error {
	user := c.Params("user")
	name := c.Params("room")

	room := j.Room(name)

	log.Println("[pages] room", room.Name)

	return c.Render("room", f.Map{
		"User": user,
		"Name": name,

		"Room": room,
	})
}
