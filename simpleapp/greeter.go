package main

import (
	"time"
)

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

type Greeter struct {
	Message Message // <- adding a Message field
	Grumpy  bool
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return NewMessage("Go away!")
	}
	return g.Message
}
