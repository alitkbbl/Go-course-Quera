package main

import "fmt"

func AddElement(numbers *[]int, element int) {
	numbers.append(element)
}
func main() {
	s := make([]int, 10)
	fmt.Println(len(s))
	s = append(s, 2)
	fmt.Println(len(s))
	AddElement(&s, 5)
	fmt.Println(s)

}
