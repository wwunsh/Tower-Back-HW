package main

import (
	"fmt"
	"sync"
)

func calculateSquare(number int, wg *sync.WaitGroup) {
	defer wg.Done()
	square := number * number
	fmt.Printf("%d\n", square)
}
func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for _, number := range numbers {
		wg.Add(1)
		go calculateSquare(number, &wg)
	}
	wg.Wait()
}
