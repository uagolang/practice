package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan int)

	go func() {
		// close channel if there will be no writes
		defer close(ch)

		n := 10
		for i := 0; i < n; i++ {
			ch <- i
		}
	}()

	count := 4
	// one channel splits to {count} channels
	// that gives more availability because if
	// your input channel blocks - data processing also blocks
	// but if we have few channels that working concurrently
	// only one of {count} channels will be blocked
	channels := split(ch, count)

	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for value := range channels[i] {
				fmt.Printf("[ch #%d] value: %d\n", i, value)
				time.Sleep(200 * time.Millisecond)
			}
		}()
	}

	wg.Wait()

	fmt.Printf("done in %s\n", time.Now().Sub(start))
}

func split[T any](inputCh <-chan T, n int) []<-chan T {
	outputChs := make([]chan T, n)
	for i := 0; i < n; i++ {
		outputChs[i] = make(chan T)
	}

	go func() {
		var index int
		for value := range inputCh {
			outputChs[index] <- value
			index = (index + 1) % n
		}

		for _, ch := range outputChs {
			close(ch)
		}
	}()

	// this block is needed only if you have
	// directed channels (<-chan T) in result
	// because golang cannot cast
	// ([]chan T) to ([]<-chan T) automatically
	resultChs := make([]<-chan T, n)
	for i := range outputChs {
		resultChs[i] = outputChs[i]
	}

	return resultChs
}
