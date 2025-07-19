package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	coats := strings.Fields(scanner.Text())[1:]
	scanner.Scan()
	shirts := strings.Fields(scanner.Text())[1:]
	scanner.Scan()
	pants := strings.Fields(scanner.Text())[1:]
	scanner.Scan()
	caps := strings.Fields(scanner.Text())[1:]
	scanner.Scan()
	jackets := strings.Fields(scanner.Text())[1:]
	scanner.Scan()
	season := scanner.Text()

	filteredCoats := []string{}
	filteredJackets := []string{}
	filteredCaps := []string{}
	switch season {
	case "SPRING":
		filteredCoats = coats
		filteredJackets = []string{}
		filteredCaps = append(caps, "")
	case "SUMMER":
		filteredCoats = []string{}
		filteredJackets = []string{}
		filteredCaps = caps
	case "FALL":
		for _, c := range coats {
			if c != "yellow" && c != "orange" {
				filteredCoats = append(filteredCoats, c)
			}
		}
		filteredJackets = []string{}
		filteredCaps = append(caps, "")
	case "WINTER":
		filteredCoats = coats
		filteredJackets = jackets
		filteredCaps = []string{}
	}

	printCombination := func(coat, shirt, pant, cap, jacket string) {
		out := []string{}
		if coat != "" {
			out = append(out, "COAT: "+coat)
		}
		if shirt != "" {
			out = append(out, "SHIRT: "+shirt)
		}
		if pant != "" {
			out = append(out, "PANTS: "+pant)
		}
		if cap != "" {
			out = append(out, "CAP: "+cap)
		}
		if jacket != "" {
			out = append(out, "JACKET: "+jacket)
		}
		fmt.Println(strings.Join(out, " "))
	}

	if season == "WINTER" {
		for _, shirt := range shirts {
			for _, pant := range pants {
				for _, coat := range filteredCoats {
					printCombination(coat, shirt, pant, "", "")
				}
				for _, jacket := range filteredJackets {
					printCombination("", shirt, pant, "", jacket)
				}
			}
		}
	} else {
		for _, shirt := range shirts {
			for _, pant := range pants {
				if len(filteredCoats) > 0 {
					for _, coat := range filteredCoats {
						for _, cap := range filteredCaps {
							if cap == "" {
								printCombination(coat, shirt, pant, "", "")
							} else {
								printCombination(coat, shirt, pant, cap, "")
							}
						}
					}
				}
				for _, cap := range filteredCaps {
					if cap == "" {
						printCombination("", shirt, pant, "", "")
					} else {
						printCombination("", shirt, pant, cap, "")
					}
				}
			}
		}
	}
}
