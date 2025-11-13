package main

import "fmt"

func main() {
	fmt.Println("------------------")
	fmt.Println("Function in golang")
	fmt.Println()

	// Simple
	result := sum(2, 3)
	fmt.Printf("%d + %d = %d", 2, 3, result)

	fmt.Println()

	// Multiple returns

	// Named return values

	// Variadic functions

	// Anonymous functions

	// Functions closure

	//

	// Built-in Functions

	fmt.Println()
	fmt.Println("------------------")
}

// Sum
func sum(a int, b int) int {
	return a + b
}
