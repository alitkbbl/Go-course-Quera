package main

import "fmt"

func handlePanic() {
	r := recover()
	if r != nil {
		fmt.Println("Panic Recovered")
	}
}
func division(a, b int) float64 {
	defer handlePanic()
	if b == 0 {
		panic("Division by zero")
	}
	return float64(a) / float64(b)
}
func main() {
	a := 3
	b := 0
	division(a, b)
	fmt.Println("End")
}
