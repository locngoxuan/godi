package main

import (
	"log"
	"time"
)

type IGreeter interface {
	Greet() IMessage
}

func NewGreeter() IGreeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return &Greeter{Grumpy: grumpy}
}

type Greeter struct {
	Grumpy bool
}

func (g Greeter) Greet() IMessage {
	if g.Grumpy {
		return NewMessage("Go away!")
	}
	var msg IMessage
	err := container.Invoke(func(m IMessage) {
		msg = m
	})
	if err != nil {
		log.Fatalln(err)
	}
	return msg
}
