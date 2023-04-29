package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go g1(ctx, ch1)
	go g1(ctx, ch2)
	go g1(ctx, ch3)

	select {
	case msg1 := <-ch1:
		fmt.Println("received", msg1)
		cancelFunc()
		close(ch1)
		close(ch2)
		close(ch3)
	case msg2 := <-ch2:
		fmt.Println("received", msg2)
		cancelFunc()
		close(ch1)
		close(ch2)
		close(ch3)
	case msg3 := <-ch3:
		fmt.Println("received", msg3)
		cancelFunc()
		close(ch1)
		close(ch2)
		close(ch3)
	}
	time.Sleep(3 * time.Second)
}

func g1(ctx context.Context, ch chan string) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Was canceled")
			return
		default:
			resp, err := http.Get("https://random-word-api.herokuapp.com/word")
			if err != nil {
				log.Fatalln(err)
			}
			text, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			resp.Body.Close()
			if ctx.Err() != nil {
				fmt.Println("Was canceled")
				return
			}
			ch <- string(text)[1 : len(text)-1]
		}
	}
}
