package main

import (
	"fmt"
	"log"
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

	eventsQueue := make([]chan SomeEvent, len(listeners))

	for i := 0; i < len(listeners); i++ {
		eventsQueue[i] = make(chan SomeEvent, 2)
	}

	for i, listener := range listeners {
		go func(i int, listener Listener) {
			for {
				event := <-eventsQueue[i]
				listener.Handle(event)
			}
		}(i, listener)
	}

	for event := range events {
		for i, queue := range eventsQueue {
			if len(queue) == 2 {
				log.Println("Message is lost from", i)
				continue
			}
			queue <- event
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
	fmt.Println(time.Now(), "Hello")
}

type listener2 struct {
}

func (l listener2) Handle(e SomeEvent) {
	fmt.Println("World")
}
