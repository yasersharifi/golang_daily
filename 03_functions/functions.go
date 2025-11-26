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

	fmt.Println()

	// Variadic functions
	variadicSumResult := variadicSum(2, 3, 4, 5, 6, 8, 9)
	fmt.Printf("Variadic Sum Result: %d", variadicSumResult)
	fmt.Println()

	log("Error", "failed to initial project", "failed to connect DB")

	// Anonymous functions

	// Functions closure

	// Higher-order functions (passing funcs as args, returning funcs)

	// Methods (value vs pointer receivers)

	// Interfaces with function behaviors

	// Error handling patterns (wrapping, sentinel errors)

	// Generics and type parameters (Go 1.18+)

	// Defer, panic, recover inside functions

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
		return 0, fmt.Errorf("b can't be zero")
	}
	return a / b, nil

}

// Named return value
func sumMines(a int, b int) (sum int, mines int) {
	s := a + b
	m := a - b

	return s, m
}

// Variadic function
func variadicSum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func log(level string, messages ...string) {
	for _, msg := range messages {
		fmt.Printf("[%s] - %s", level, msg)
		fmt.Println()
	}
}
