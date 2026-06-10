package main

import (
	"fmt"
	"sync"
)

func pow(i int, wg *sync.WaitGroup, channel chan int) {
	var square int

	square = i * i
	fmt.Println("Square of", i, "is", square)
	channel <- square
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var channel chan int = make(chan int)
	var sum int
	ilist := []int{1, 2, 3, 4, 5}

	for _, i := range ilist {
		wg.Add(1)
		go pow(i, &wg, channel)
	}

	for j := 0; j < len(ilist); j++ {
		sum += <-channel
	}

	wg.Wait()
	fmt.Println("sum", sum)
}
