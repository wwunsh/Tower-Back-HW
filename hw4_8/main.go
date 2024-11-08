package main

import "fmt"

func array(arr []int, ch chan int) {
	for _, v := range arr {
		ch <- v
	}
	close(ch)
}
func double(arr []int, ch chan int) {
	for _, num := range arr {
		ch <- num * 2
	}
	close(ch)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	ch := make(chan int)
	ch2 := make(chan int)

	go array(arr, ch)

	for value := range ch {
		fmt.Println(value)
	}

	go double(arr, ch2)

	for product := range ch2 {
		fmt.Println(product)
	}

}
