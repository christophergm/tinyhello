package main

import (
	"fmt"
)

func calculatePi(iterations int) float64 {
	pi := 0.0
	sign := 1.0

	for i := 0; i < iterations; i++ {
		term := sign / float64(2*i + 1)
		pi += term
		sign = -sign // flip sign for the next term
	}
	return pi * 4.0
}
func main() {
	numIterations := 1000000000 // numbers of iterations for approx

	approxPi := calculatePi(numIterations)

	fmt.Printf("approx of pi using leibiz formula: %.30f",approxPi)
	
}
