package main

import (
	"fmt"
	"sync"
)

func square(number int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	square := number * number
	ch <- square
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	sum := 0
	var wg sync.WaitGroup
	ch := make(chan int, len(arr))
	for _, number := range arr {
		wg.Add(1)
		go square(number, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	for sq := range ch {
		sum += sq
	}
	fmt.Printf("%d\n", sum)
}
