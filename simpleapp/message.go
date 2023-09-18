package main

type Message struct {
	v string
}

func (m Message) Value() string {
	return m.v
}

func NewMessage(phrase string) Message {
	return Message{
		v: phrase,
	}
}
