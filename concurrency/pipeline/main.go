package main

import (
	"fmt"
)

func main() {
	// integer pipeline
	ints := generate(1, 2, 3, 4, 5)
	appliedInts := apply(ints, func(n int) int { return n * n })
	aggregatedInts := aggregate(appliedInts)

	fmt.Println("sum of squares:", <-aggregatedInts)

	// float pipeline
	floats := generate(1.5, 2.5, 3.5)
	appliedFloats := apply(floats, func(f float64) float64 { return f * 2 })
	aggregatedFloats := aggregate(appliedFloats)

	fmt.Println("sum of float values:", <-aggregatedFloats)
}

func generate[T any](values ...T) <-chan T {
	resultCh := make(chan T)

	go func() {
		defer close(resultCh)

		for _, v := range values {
			resultCh <- v
		}
	}()

	return resultCh
}

func apply[T any, R any](in <-chan T, transform func(T) R) <-chan R {
	resultCh := make(chan R)

	go func() {
		defer close(resultCh)

		for v := range in {
			resultCh <- transform(v)
		}
	}()

	return resultCh
}

func aggregate[T int | float64](in <-chan T) <-chan T {
	resultCh := make(chan T)

	go func() {
		defer close(resultCh)

		var sum T
		for v := range in {
			sum += v
		}

		resultCh <- sum
	}()

	return resultCh
}
