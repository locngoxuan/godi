package main

import "log"

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

type Event struct {
	Greeter Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	log.Println(msg.Value())
}
