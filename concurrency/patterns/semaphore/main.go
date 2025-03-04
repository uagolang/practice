package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	data := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	// buffer size is 1, so
	// only 1 goroutine can work concurrently in a moment of time
	channelSemaphore(1, data)
	structSemaphore(1, data)

	// buffer size n (n > 1), so
	// only n goroutines can work concurrently in a moment of time
	channelSemaphore(5, data)
	structSemaphore(5, data)
}

func structSemaphore(size int, data []string) {
	if size < 1 {
		size = 1
	}

	start := time.Now()
	sem := NewSemaphore(size)

	var wg sync.WaitGroup
	for i := 0; i < len(data); i++ {
		wg.Add(1)
		sem.Acquire() // block channel to do some work

		go func(data []string, i int) {
			// unblock channel after return
			defer func() {
				// need to unblock channel first
				sem.Release()
				wg.Done()
			}()

			// doing work
			println("process:", data[i])
			time.Sleep(500 * time.Millisecond)

			//sem.Release() // or you can unblock it manually at the end
			//wg.Done()
		}(data, i)
	}

	wg.Wait()

	fmt.Printf("SUCCESS: struct semaphore with buffer %d done in %s\n\n", size, time.Now().Sub(start))
}

func channelSemaphore(size int, data []string) {
	if size < 1 {
		size = 1
	}

	start := time.Now()
	sem := make(chan struct{}, size)

	var wg sync.WaitGroup
	for i := 0; i < len(data); i++ {
		wg.Add(1)
		go func(data []string, i int) {
			// unblock channel after return
			defer func() {
				// need to unblock channel first
				<-sem
				wg.Done()
			}()
			sem <- struct{}{} // block channel to do some work

			// doing work
			println("process:", data[i])
			time.Sleep(500 * time.Millisecond)

			//<-sem // or you can unblock it manually at the end
			//wg.Done()
		}(data, i)
	}

	wg.Wait()
	close(sem)

	fmt.Printf("SUCCESS: semaphore with buffer %d done in %s\n\n", size, time.Now().Sub(start))
}

type Semaphore struct {
	size int
	ch   chan struct{}
}

func NewSemaphore(size int) *Semaphore {
	return &Semaphore{size: size, ch: make(chan struct{}, size)}
}

func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.ch
}
