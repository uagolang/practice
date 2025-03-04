package main

import (
	"fmt"
	"sync"
)

// you can ask: we have fan-out pattern, why do we need this?

// answer is simple:
// in fan-out data processes concurrently
// but in tee pattern all data copies to returned channels

func main() {
	ch := make(chan int)

	// writing data to the channel
	go func() {
		defer close(ch)

		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// split 1 channel to 2 with same data
	ch1, ch2 := tee(ch)

	var wg sync.WaitGroup
	wg.Add(2)

	// read channel 1
	go func() {
		defer wg.Done()

		for v := range ch1 {
			fmt.Println("ch1 value:", v)
		}
	}()

	// read channel 2
	go func() {
		defer wg.Done()

		for v := range ch2 {
			fmt.Println("ch2 value:", v)
		}
	}()

	wg.Wait()
	fmt.Println("done")
}

// super easy logic - just create 2 channels
// and set data to each from input channel
func tee[T any](ch chan T) (chan T, chan T) {
	ch1, ch2 := make(chan T), make(chan T)

	go func() {
		defer func() {
			close(ch1)
			close(ch2)
		}()

		for data := range ch {
			ch1 <- data
			ch2 <- data
		}
	}()

	return ch1, ch2
}
