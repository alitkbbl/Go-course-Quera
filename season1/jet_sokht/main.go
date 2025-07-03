package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	names := make([]string, n)
	counts := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		names[i] = parts[0]

		valsStr := parts[1:]
		vals := make([]int, len(valsStr))
		for j := 0; j < len(valsStr); j++ {
			vals[j], _ = strconv.Atoi(valsStr[j])
		}

		counts[i] = countArithmeticSubarrays(vals)
	}

	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}

	sort.Slice(indices, func(i, j int) bool {
		if counts[indices[i]] == counts[indices[j]] {
			return names[indices[i]] < names[indices[j]]
		}
		return counts[indices[i]] > counts[indices[j]]
	})

	for _, idx := range indices {
		fmt.Println(names[idx], counts[idx])
	}
}

func countArithmeticSubarrays(arr []int) int {
	count := 0
	for start := 0; start < len(arr)-2; start++ {
		diff := arr[start+1] - arr[start]
		for end := start + 2; end < len(arr); end++ {
			valid := true
			for k := start; k < end; k++ {
				if arr[k+1]-arr[k] != diff {
					valid = false
					break
				}
			}
			if valid {
				count++
			} else {
				break
			}
		}
	}
	return count
}
