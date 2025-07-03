package main

import "fmt"

func main() {
	var n, sum uint64
	fmt.Scanf("%d", &n)
	if n <= 100 {
		sum = n / 20
	} else if n <= 500 {
		sum = 5 + (n-100)/10
	} else if n <= 1000 {
		sum = 5 + 40 + (n-500)*15/100
	} else {
		sum = 5 + 40 + 75 + (n-1000)/5
	}
	fmt.Println(sum)
}
