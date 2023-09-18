package main

import (
	"log"
	"time"
)

type IGreeter interface {
	Greet() IMessage
}

func NewGreeter() IGreeter {
	var msg IMessage
	err := container.Invoke(func(m IMessage) {
		msg = m
	})
	if err != nil {
		log.Fatalln(err)
	}
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return &Greeter{Message: msg, Grumpy: grumpy}
}

type Greeter struct {
	Message IMessage // <- adding a Message field
	Grumpy  bool
}

func (g Greeter) Greet() IMessage {
	if g.Grumpy {
		return NewMessage("Go away!")
	}
	return g.Message
}
