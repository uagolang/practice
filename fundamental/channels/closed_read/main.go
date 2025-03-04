package main

import (
	"fmt"
)

type test struct {
	name string
}

func main() {
	ch := make(chan *test)

	go func() {
		ch <- &test{name: "test"}
	}()

	// output here will be 'test' as expected
	fmt.Printf("channel value before close: %v\n", <-ch)

	go func() {
		close(ch)
	}()

	// if channel has been closed already, there will be no panic
	// but will be returned nil value of channel type
	fmt.Printf("channel value from closed channel: %v\n", <-ch)
}
