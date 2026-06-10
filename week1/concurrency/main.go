package main

import (
	"fmt"
	"sync"
)

func pow(i int, wg *sync.WaitGroup) {
	fmt.Println("Square of", i, "is", i*i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	ilist := []int{1, 2, 3, 4, 5}
	for _, i := range ilist {
		wg.Add(1)
		go pow(i, &wg)
	}

	wg.Wait()
}
