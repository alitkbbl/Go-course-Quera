package main

import "fmt"

func ConvertToDigitalFormat(hour, minute, second int) string {
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func ExtractTimeUnits(seconds int) (int, int, int) {
	hour := seconds / 3600
	remainder := seconds % 3600
	minute := remainder / 60
	second := remainder % 60
	return hour, minute, second
}
