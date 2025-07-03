package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	scanner := bufio.NewScanner(os.Stdin)
	codeMap := make(map[string]string)

	for i := 0; i < n; i++ {
		scanner.Scan()
		inputCountryCode := scanner.Text()
		parts := strings.Fields(inputCountryCode)
		codeMap[parts[1]] = parts[0]
	}

	var q int
	fmt.Scanf("%d\n", &q)

	var myAnswer []string

	for i := 0; i < q; i++ {
		scanner.Scan()
		phoneNumberInput := scanner.Text()

		if len(phoneNumberInput) < 3 {
			myAnswer = append(myAnswer, "Invalid Number")
			continue
		}
		myKey := phoneNumberInput[:3]

		country, ok := codeMap[myKey]
		if ok {
			myAnswer = append(myAnswer, country)
		} else {
			myAnswer = append(myAnswer, "Invalid Number")
		}
	}

	for i := 0; i < len(myAnswer); i++ {
		fmt.Println(myAnswer[i])
	}
}
