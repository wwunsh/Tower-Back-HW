package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map
	var wg sync.WaitGroup

	numGoroutines := 10

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sm.Store(id, id*id)
			fmt.Printf("Goroutine %d stored: %d -> %d\n", id, id, id*id)
		}(i)
	}

	wg.Wait()

	for i := 0; i < numGoroutines; i++ {
		if value, ok := sm.Load(i); ok {
			fmt.Printf("Key: %d, Value: %d\n", i, value)
		} else {
			fmt.Printf("Key %d not found\n", i)
		}
	}
}
