package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]int{
		"gold":   1,
		"silver": 2,
		"bronze": 3,
	}
	bytes, err := json.Marshal(m)
	if err != nil {
		panic("failed to marshal")
	}
	fmt.Println(bytes)
}
