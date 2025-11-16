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
	division(15, 3)

	// Named return values
	sumMines(4, 2)

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

// Multiple returns
func division(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("b can't be zero");
	}
	return a / b, nil
	
}

// Named return value
func sumMines(a int, b int) (sum int, mines int) {
	s := a + b
	m := a - b

	return s, m
}
