package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	generator := func(done <-chan interface{}, integers ...float64) <-chan float64 {
		intStream := make(chan float64)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func(
		done <-chan interface{},
		intStream <-chan float64,
		multiplier float64,
	) <-chan float64 {
		multipliedStream := make(chan float64)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	add := func(
		done <-chan interface{},
		intStream <-chan float64,
		additive float64,
	) <-chan float64 {
		addedStream := make(chan float64)
		go func() {
			defer close(addedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:
				}
			}
		}()
		return addedStream
	}

	done := make(chan interface{})
	defer close(done)

	n := 1000000

	// Generate a slice of random integers
	var ints []float64
	for i := 1; i < n; i++ {
		ints = append(ints, rand.Float64()*float64(n))
	}

	intStream := generator(done, ints...)

	// Measure the time taken to multiply and add
	start := time.Now()
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	elapsed := time.Since(start)

	for v := range pipeline {
		fmt.Println(v)
	}

	fmt.Printf("Time taken: %s\n", elapsed)
}
