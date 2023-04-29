package main

import (
	"fmt"
	"time"
)

func main() {
	events := make(chan SomeEvent)

	listeners := []Listener{listener1{}, listener2{}}

	go func() {
		for {
			events <- SomeEvent{}
			time.Sleep(1 * time.Second)
		}
	}()

	for event := range events {
		for _, listener := range listeners {
			listener.Handle(event)
		}
	}
}

type SomeEvent struct {
}

type Listener interface {
	Handle(e SomeEvent)
}

type listener1 struct {
}

func (l listener1) Handle(e SomeEvent) {
	time.Sleep(2 * time.Second)
	fmt.Println("Hello")
}

type listener2 struct {
}

func (l listener2) Handle(e SomeEvent) {
	fmt.Println("World")
}
