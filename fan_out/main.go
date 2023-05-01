package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		oscall := <-quit
		log.Printf("system call:%+v", oscall)
		cancelFunc()
	}()

	events := make(chan SomeEvent)
	listeners := []Listener{listener1{}, listener2{}}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				log.Println("Canceling 1")
				return
			default:
				events <- SomeEvent{}
				time.Sleep(1 * time.Second)
			}
		}
	}()

	eventsQueue := make([]chan SomeEvent, len(listeners))

	for i := 0; i < len(listeners); i++ {
		eventsQueue[i] = make(chan SomeEvent, 2)
	}

	for i, listener := range listeners {
		wg.Add(1)
		go func(i int, listener Listener) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					log.Println("Canceling 2")
					return
				default:
					event := <-eventsQueue[i]
					listener.Handle(event)
				}
			}
		}(i, listener)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for event := range events {
			for i, queue := range eventsQueue {
				select {
				case <-ctx.Done():
					log.Println("Canceling 3")
					return
				case queue <- event:
				default:
					log.Println("Message is lost from", i)
				}
			}
		}
	}()

	wg.Wait()
	fmt.Println("Main done")
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
