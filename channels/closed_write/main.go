package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 0
	}()

	// output here will be 'test' as expected
	fmt.Printf("channel value before close: %v\n", <-ch)

	go func() {
		close(ch)
	}()

	// will be panic here
	// golang doesn't allow writing to the closed channel
	// so you should remember a rule:
	// writer should close the channel
	ch <- 1
}
