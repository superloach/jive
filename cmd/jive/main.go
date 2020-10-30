package main

import (
	"flag"

	"git.superloach.xyz/_/jive"
)

var (
	addr  = flag.String("addr", ":7173", "host:port to bind on")
	www   = flag.String("www", "./www", "path for web files")
	room1 = flag.String("room1", "lobby", "first room to create")
)

func main() {
	flag.Parse()

	err := jive.New(jive.Opts{
		Addr:  *addr,
		WWW:   *www,
		Room1: *room1,
	}).Serve()
	if err != nil {
		panic(err)
	}
}
