package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isArmstrong(num int) bool {
	strNum := strconv.Itoa(num)
	n := len(strNum)
	sum := 0
	for _, ch := range strNum {
		digit := int(ch - '0')
		sum += int(math.Pow(float64(digit), float64(n)))
	}
	return sum == num
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if len(input) == 0 {
		fmt.Println("YES")
		return
	}

	var numbers []string
	var currentNum strings.Builder

	for _, ch := range input {
		if unicode.IsDigit(ch) {
			currentNum.WriteRune(ch)
		} else {
			if currentNum.Len() > 0 {
				numbers = append(numbers, currentNum.String())
				currentNum.Reset()
			}
		}
	}
	if currentNum.Len() > 0 {
		numbers = append(numbers, currentNum.String())
	}

	sum := 0
	for _, strNum := range numbers {
		num, err := strconv.Atoi(strNum)
		if err == nil {
			sum += num
		}
	}

	if isArmstrong(sum) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
