package main

import (
	"fmt"
	"math"
)

func calculateSquareRoot(num float64) float64 {
	if num < 0 {
		panic("cannot calculate square root of a negative number")
	}
	return math.Sqrt(num)
}
func main() {
	number := -5.0
	result := calculateSquareRoot(number)
	fmt.Println(result)
}
