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

	// Infinite Loops - while (simulate) - while(1)
	j := 0
	for {
		fmt.Println("J is: ", j)
		j += 1
		if j > 20 {
			break
		}
	}

	// for-range
	slice := [3]int{1, 3, 5}
	for i, s := range slice {
		fmt.Printf("Slice item[%d] is: %d", i, s)
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("End of loops in golang")
	fmt.Println("-----------------")
}
