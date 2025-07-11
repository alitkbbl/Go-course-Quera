package main

import (
	"errors"
	"fmt"
)

func main() {
	rootErr := errors.New("Root error")
	// Wrap the root error with additional context
	wrappedErr1 := fmt.Errorf("Error 1: %w", rootErr)
	wrappedErr2 := fmt.Errorf("Error 2: %w", wrappedErr1)
	fmt.Printf("wrappedErr2: %v\n", wrappedErr2)

	fmt.Printf("wrappedErr2: %v\n", wrappedErr2)
	fmt.Printf("Unwrap(wrappedErr2): %v\n", errors.Unwrap(wrappedErr2))
	fmt.Printf("Unwrap(Unwrap(wrappedErr2)): %v\n", errors.Unwrap(errors.Unwrap(wrappedErr2)))
}
