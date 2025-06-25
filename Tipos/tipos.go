package main

import (
	"fmt"
)

func main() {
	// Booleano
	fmt.Printf("Type: %T - Value: %v\n", true, true)

	// String
	fmt.Printf("Type: %T - Value: %v\n", "steph", "steph")
	fmt.Printf("Type: %T - Value: %v\n", "1", "1")

	// Int
	fmt.Printf("Type: %T - Value: %v\n", 1, 1)

	// Float (float64/float32)
	fmt.Printf("Type: %T - Value: %v\n", 1.233, 1.233)
}
