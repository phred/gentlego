package main

import (
	"fmt"
	"time"
)

func counter(response chan int) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		response <- i
	}
	close(response)
}

func main() {
	ch := make(chan int)

	go counter(ch)
	for resp := range(ch) {
		fmt.Printf("Got: %d!\n", resp)
	}
}
