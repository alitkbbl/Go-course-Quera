package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	userName := make([]string, 0, n)
	scoreStatus := make([]string, 0, n)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		scanner.Scan()
		inputUserName := scanner.Text()
		userName = append(userName, inputUserName)

		scanner.Scan()
		scoresLine := scanner.Text()
		scoresStr := strings.Fields(scoresLine)

		sum := 0
		for _, val := range scoresStr {
			score, err := strconv.Atoi(val)
			if err == nil {
				sum += score
			}
		}
		avg := sum / len(scoresStr) //avg

		if avg >= 80 {
			scoreStatus = append(scoreStatus, "Excellent")
		} else if avg >= 60 {
			scoreStatus = append(scoreStatus, "Very Good")
		} else if avg >= 40 {
			scoreStatus = append(scoreStatus, "Good")
		} else {
			scoreStatus = append(scoreStatus, "Fair")
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s %s\n", userName[i], scoreStatus[i])
	}
}
