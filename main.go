package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	const epsilon = 1e-10
	for i := 0; i < 10; i++{
		prev := z
		z -= (z*z - x) / (2*z)
		if math.Abs(z-prev) < epsilon { 
			return z
		}
		fmt.Printf("%g\n", z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
