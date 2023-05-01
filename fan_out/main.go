package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	go func() {
		http.ListenAndServe(":8080", nil)
	}()
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)

	events := make(chan SomeEvent)
	listeners := []Listener{listener1{}, listener2{}}
	var wg sync.WaitGroup

	//wg.Add(1)
	go func() {
		//defer wg.Done()
		tick := time.Tick(time.Second)
		for {
			select {
			case <-ctx.Done():
				log.Println("Canceling event maker")
				close(events)
				return
			case <-tick:
				events <- SomeEvent{}
			}
		}
	}()

	eventsQueue := make([]chan SomeEvent, len(listeners))
	for i := 0; i < len(listeners); i++ {
		eventsQueue[i] = make(chan SomeEvent, 2)
	}

	//wg.Add(1)
	go func() {
		//defer wg.Done()
		for event := range events {
			for i, queue := range eventsQueue {
				select {
				case <-ctx.Done():
					close(queue)
					log.Println("Canceling 2")
					return
				case queue <- event:
				default:
					log.Println("Message is lost from", i)
				}
			}
		}
		log.Println("Canceling 3")
	}()

	for i, listener := range listeners {
		wg.Add(1)
		go func(i int, listener Listener) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					log.Println("Canceling 3", i)
					return
				case event, ok := <-eventsQueue[i]:
					if !ok {
						log.Println("closed 3", i)
						break
					}
					listener.Handle(event)
				}
			}
		}(i, listener)
	}

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
	time.Sleep(5 * time.Second)
	fmt.Println(time.Now(), "Hello")
}

type listener2 struct {
}

func (l listener2) Handle(e SomeEvent) {
	fmt.Println(time.Now(), "World")
}
