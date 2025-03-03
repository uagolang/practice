package main

import (
	"fmt"
	"sync"
	"time"
)

// main idea of that pattern is to merge
// list of channels into resulting one

func main() {
	start := time.Now()

	// init for example 4 channels
	ch1, ch2, ch3, ch4 := make(chan int), make(chan int), make(chan int), make(chan int)

	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			close(ch3)
			close(ch4)
		}()

		for i := 0; i < 10; i++ {
			ch1 <- i
			ch2 <- i * 5
			ch3 <- i * 6
			ch4 <- i * 7
		}
	}()

	// function merge returns channel
	// and we can iterate over the channel
	for value := range merge(ch1, ch2, ch3, ch4) {
		fmt.Println("value:", value)
	}

	fmt.Println("done in", time.Now().Sub(start))
}

func merge[T any](chs ...chan T) chan T {
	// we need wait group to sync goroutines
	// and close result channel
	var wg sync.WaitGroup
	wg.Add(len(chs))

	result := make(chan T)
	for _, ch := range chs {
		// wg.Add(1) could be here instead of line 47
		go func() {
			defer wg.Done()

			for value := range ch {
				result <- value
			}
		}()
	}

	// wait all groups to be done & close channel in background
	go func() {
		wg.Wait()
		close(result)
	}()

	// returns channel immediately
	return result
}
