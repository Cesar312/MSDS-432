package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
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
		intStream <-chan int,
		multiplier int,
	) <-chan int {
		multipliedStream := make(chan int)
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
		intStream <-chan int,
		additive int,
	) <-chan int {
		addedStream := make(chan int)
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

	n := 10000

	// Generate a slice of random integers
	var ints []int
	for i := 1; i < n; i++ {
		ints = append(ints, rand.Intn(1000))
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
