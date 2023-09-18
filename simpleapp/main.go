package main

func main() {
	msg := NewMessage("Hi there!")
	greeter := NewGreeter(msg)
	event := NewEvent(greeter)
	event.Start()
}
