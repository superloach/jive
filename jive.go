package jive

import (
	"log"
	"sync"
	"time"

	f "github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	ws "github.com/gofiber/websocket/v2"
)

type Jive struct {
	sync.Mutex

	Opts Opts

	Rooms  []*Room
	Spells map[string]Spell
}

func New(o Opts) *Jive {
	log.Println("[jive] new")

	return &Jive{
		Opts: o,
	}
}

func (j *Jive) Serve() error {
	j.Spells = map[string]Spell{
		"test": TestSpell,
	}

	_ = j.Room(j.Opts.Room1)

	app := f.New(f.Config{
		Views:        html.New(j.Opts.WWW, ".html"),
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
	})

	app.Get("/", j.IndexPage)

	app.Get("/favicon.ico", func(*f.Ctx) error {
		return f.NewError(f.StatusNotFound)
	})

	app.Static("/", j.Opts.WWW)

	app.Get("/:room", j.RoomPage)

	app.Get("/talk/:user/:room", ws.New(j.Talk))

	log.Println("[jive] serve", j.Opts.Addr)
	return app.Listen(j.Opts.Addr)
}
