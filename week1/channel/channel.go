package main

import "fmt"

func pow(i int, channel chan int) {
	square := i * i
	fmt.Println("Square of", i, "is", square)
	channel <- square
}

func main() {
	channel := make(chan int)
	var sum int
	ilist := []int{1, 2, 3, 4, 5}

	for _, i := range ilist {
		go pow(i, channel)
	}

	for j := 0; j < len(ilist); j++ {
		sum += <-channel
	}

	fmt.Println("sum:", sum)
}
