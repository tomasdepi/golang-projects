package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var LOOP_COUNT int = 100

func calculateSquare(i int) {
	defer wg.Done()
	fmt.Println(i * i)
}

func main() {
	start := time.Now()
	wg.Add(LOOP_COUNT)

	for i := 1; i <= LOOP_COUNT; i++ {
		go calculateSquare(i)
	}

	wg.Wait()

	finish := time.Since(start)

	fmt.Println("Function took ", finish)
}
