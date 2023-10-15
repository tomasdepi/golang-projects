package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) // channel of type string, unbuffered
	go sell(ch, "Hellooooo")
	go buy(ch)
	time.Sleep(2 * time.Second)
}

// sends data to channel
func sell(ch chan string, m string) {
	ch <- m // this blocks until someone (another go-rutine) reads from channel
	fmt.Println("Sending data to channel")
}

// reads data from channel
func buy(ch chan string) {
	fmt.Println("Waiting for data.....")
	val := <-ch // blocks the execution util someone writes to channel
	fmt.Println("Data received: ", val)
}
