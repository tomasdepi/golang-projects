// a buffered channel has capacity to hold data
// sending to a channel blocks the go rutine only if the channel is full
// reading from a channel blocks the go rutine only if the channel is empty

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 3)

	wg.Add(2)
	go sell(ch)
	wg.Wait()
}

func sell(ch chan int) {
	ch <- 10
	ch <- 11
	ch <- 12

	go buy(ch)

	fmt.Println("Sent three values to channel")
	close(ch)
	wg.Done()
}

func buy(ch chan int) {
	fmt.Println("Waiting for data")

	for val := range ch { // this works due the channel is closed, otherwise produces deadlock
		fmt.Println("Received data: ", val)
	}

	wg.Done()
}
