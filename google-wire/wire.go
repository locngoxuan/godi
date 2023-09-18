//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// wire_gen.go
func InitializeEvent(phrase string) IEvent {
	wire.Build(NewEvent,
		NewGreeter,
		NewMessage)
	return &Event{}
}
