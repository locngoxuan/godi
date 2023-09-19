package main

import (
	"log"
)

func main() {
	Initialize()
	err := container.Invoke(func(e IEvent) {
		e.Start()
	})
	if err != nil {
		log.Fatalln(err)
	}
}
