// select statement if like switch but for channels
// select statement lets a go-rutine wait on multiple communication operations
// in select, each case statement waits for a send or receive operation from a channel
// select blocks until any of the case statement are ready
// if multiple case statement are ready, it selects one at random and proceeds

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	time.Sleep(1 * time.Second)

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case val2 := <-ch2:
		fmt.Println(val2)
	default: // in case any of the case is ready, instead of blocking executes de default
		fmt.Println("Default case")
	}

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case <-time.After(2 * time.Second): // also is possible to use the time.After method to timeout
		fmt.Println("Timeout Case")
	}

}

func goOne(ch chan string) {
	ch <- "Channel 1"
}

func goTwo(ch chan string) {
	ch <- "Channel 2"
}
