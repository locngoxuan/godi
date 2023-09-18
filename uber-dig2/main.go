package main

import (
	"log"

	"go.uber.org/dig"
)

var container = dig.New()

func main() {
	Initialize()
	err := container.Invoke(func(e IEvent) {
		e.Start()
	})
	if err != nil {
		log.Fatalln(err)
	}
}
