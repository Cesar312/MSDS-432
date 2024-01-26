package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	multiply := func(value, multiplier float64) float64 {
		return value * multiplier
	}

	add := func(value, additive float64) float64 {
		return value + additive
	}

	n := 1000000

	// Generate a slice of n random float64 numbers
	floats := make([]float64, n)
	for i := range floats {
		floats[i] = rand.Float64() * float64(n) // Random float64number between 0 and 1, then scale up to n
	}
	// Measure the time taken to multiply and add
	start := time.Now()
	for _, v := range floats {
		fmt.Println(multiply(add(multiply(v, 2), 1), 2))
	}
	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)
}
