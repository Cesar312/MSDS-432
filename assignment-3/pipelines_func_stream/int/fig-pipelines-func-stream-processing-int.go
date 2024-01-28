package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	multiply := func(value, multiplier int) int {
		return value * multiplier
	}

	add := func(value, additive int) int {
		return value + additive
	}

	n := 10000

	// Generate a slice of n random integers
	ints := make([]int, n)
	for i := range ints {
		ints[i] = rand.Intn(n) + 1 // Random integers between 1 and n
	}

	// Measure the time taken to multiply and add
	start := time.Now()
	for _, v := range ints {
		fmt.Println(multiply(add(multiply(v, 2), 1), 2))
	}
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)
}
