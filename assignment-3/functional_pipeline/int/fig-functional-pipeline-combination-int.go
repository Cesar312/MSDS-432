package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Function to mulitply each element in slice by a multiplier
	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, v := range values {
			multipliedValues[i] = v * multiplier
		}
		return multipliedValues
	}
	// Function to add an additive to each element in the slice
	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}

	n := 1000000

	// Generate a slice of n random integers
	ints := make([]int, n)
	for i := range ints {
		ints[i] = rand.Intn(n) + 1 // Random integers between 1 and n
	}

	// Measure the time taken to multiply and add
	start := time.Now()
	result := add(multiply(ints, 2), 1)
	elapsed := time.Since(start)

	for _, v := range result {
		fmt.Println(v)
	}

	fmt.Printf("Time taken: %s\n", elapsed)
}
