package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, dataCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d exiting\n", id)
			return
		case data := <-dataCh:
			fmt.Printf("Worker %d %d\n", id, data)
		}
	}
}

func readNumWorkers() int {
	var numWorkers int
	fmt.Print("Enter the number")
	if _, err := fmt.Scan(&numWorkers); err != nil || numWorkers <= 0 {
		fmt.Println("Incorrect input")
		os.Exit(1)
	}
	return numWorkers
}

func startWorkers(ctx context.Context, numWorkers int, dataCh <-chan int, wg *sync.WaitGroup) {
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i+1, dataCh, wg)
	}
}

func startDataGeneration(ctx context.Context, dataCh chan<- int) {
	rand.Seed(time.Now().UnixNano())
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(dataCh)
				return
			case dataCh <- rand.Intn(100):
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
}

func handleShutdown(cancel context.CancelFunc, wg *sync.WaitGroup) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh
	fmt.Println("\nReceived interrupt signal")
	cancel()
	wg.Wait()
	fmt.Println("Exited.")
}

func main() {
	numWorkers := readNumWorkers()
	dataCh := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	startWorkers(ctx, numWorkers, dataCh, wg)
	startDataGeneration(ctx, dataCh)
	handleShutdown(cancel, wg)
}
