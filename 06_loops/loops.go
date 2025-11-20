package main

import "fmt"

func main() {
	fmt.Println("-----------------")
	fmt.Println("Loops in golang")
	fmt.Println()

	// for
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum is: ", sum)

	// while (simulate)

	// Infinite Loops

	// for-range

	fmt.Println()
	fmt.Println("End of loops in golang")
	fmt.Println("-----------------")
}
