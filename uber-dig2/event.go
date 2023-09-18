package main

import (
	"fmt"
	"log"
)

type IEvent interface {
	Start()
}

func NewEvent() IEvent {
	var greeter IGreeter
	err := container.Invoke(func(g IGreeter) {
		greeter = g
	})
	if err != nil {
		log.Fatalln(err)
	}
	return &Event{Greeter: greeter}
}

type Event struct {
	Greeter IGreeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
