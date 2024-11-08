package main

import (
	"context"
	"fmt"
	"time"
)

func sendData(ctx context.Context, ch chan<- int) {
	counter := 0
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- counter:
			counter++
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func receiveData(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-ch:
			fmt.Printf("%d\n", data)
		}
	}
}

func main() {
	var duration int
	if _, err := fmt.Scan(&duration); err != nil || duration <= 0 {
		fmt.Println("Incorrect input")
		return
	}

	dataCh := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(duration)*time.Second)
	defer cancel()

	go sendData(ctx, dataCh)
	go receiveData(ctx, dataCh)
	<-ctx.Done()
}
