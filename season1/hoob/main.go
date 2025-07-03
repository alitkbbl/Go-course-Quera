package main

import "fmt"

func main() {
	var p, q int
	fmt.Scan(&p, &q)

	for i := 1; i <= q; i++ {
		if i%p == 0 {
			x := i / p
			for j := 0; j < x; j++ {
				fmt.Printf("Hope ")
			}
			fmt.Println()
		} else {

			fmt.Println(i)
		}
	}
}
