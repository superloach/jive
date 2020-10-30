package jive

import (
	"fmt"
	"log"
	"time"
)

const (
	TypeSpl = "spl"
)

type Event struct {
	When time.Time `json:"when"`
	Type string    `json:"type"`
	User string    `json:"user"`

	Msg string   `json:"msg,omitempty"`
	Err string   `json:"err,omitempty"`
	Spl []string `json:"spl,omitempty"`
}

func (e Event) String() string {
	switch e.Type {
	case "hey":
		return fmt.Sprintf("[hey] %s", e.User)
	case "msg":
		return fmt.Sprintf("[msg] %s: %q", e.User, e.Msg)
	case "err":
		return fmt.Sprintf("[err] %s: %q", e.User, e.Err)
	case TypeSpl:
		return fmt.Sprintf("[spl] %s: %q", e.User, e.Spl)
	default:
		return fmt.Sprintf("[%s] %s: %#v", e.Type, e.User, e)
	}
}

func HeyEvent(user string) *Event {
	log.Println("[event] new hey", user)

	return &Event{
		When: time.Now(),
		Type: "hey",
		User: user,
	}
}

func MsgEvent(user, cont string) *Event {
	log.Println("[event] new msg", user, cont)

	return &Event{
		When: time.Now(),
		Type: "msg",
		User: user,
		Msg:  cont,
	}
}

func ErrEvent(user string, err interface{}) *Event {
	log.Println("[event] new err", user, err)

	return &Event{
		When: time.Now(),
		Type: "err",
		User: user,
		Err:  fmt.Sprint(err),
	}
}

func SplEvent(user, spl string, args ...string) *Event {
	log.Println("[event] new spl", user, spl)

	return &Event{
		When: time.Now(),
		Type: TypeSpl,
		User: user,
		Spl:  append([]string{spl}, args...),
	}
}
