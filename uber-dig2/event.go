package main

import (
	"log"
)

type IEvent interface {
	Start()
}

func NewEvent() IEvent {
	return &Event{}
}

type Event struct {
}

func (e Event) Start() {
	var greeter IGreeter
	err := container.Invoke(func(g IGreeter) {
		greeter = g
	})
	if err != nil {
		log.Fatalln(err)
	}
	msg := greeter.Greet()
	log.Println(msg.Value())
}
