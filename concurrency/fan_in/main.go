package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
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

	for value := range join(ch1, ch2, ch3, ch4) {
		fmt.Println("value:", value)
	}

	fmt.Println("done in", time.Now().Sub(start))
}

func join[T any](chs ...chan T) chan T {
	var wg sync.WaitGroup
	wg.Add(len(chs))

	result := make(chan T)
	for _, ch := range chs {
		go func() {
			defer wg.Done()

			for value := range ch {
				result <- value
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}
