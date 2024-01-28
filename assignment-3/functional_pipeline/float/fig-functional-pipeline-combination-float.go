package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Function to mulitply each element in slice by a multiplier
	multiply := func(values []float64, multiplier float64) []float64 {
		multipliedValues := make([]float64, len(values))
		for i, v := range values {
			multipliedValues[i] = v * multiplier
		}
		return multipliedValues
	}
	// Function to add an additive to each element in the slice
	add := func(values []float64, additive float64) []float64 {
		addedValues := make([]float64, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}

	n := 10000

	// Generate a slice of n random float64 numbers
	floats := make([]float64, n)
	for i := range floats {
		// Generate random float64number between 0 and 1, then scale up to n
		floats[i] = rand.Float64() * float64(n)
	}

	// Measure the time taken to multiply and add
	start := time.Now()
	result := add(multiply(floats, 2), 1)
	elapsed := time.Since(start)

	for _, v := range result {
		fmt.Println(v)
	}

	fmt.Printf("Time taken: %s\n", elapsed)
}
