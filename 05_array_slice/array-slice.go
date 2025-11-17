package main

import "fmt"

func main() {
	fmt.Println("--------------")
	fmt.Println("Array and slice in golang")
	fmt.Println()

	// Define an array
	arr := [4]int{2, 4, 6, 8}
	fmt.Println("arr: ", arr)
	fmt.Println("first item: ", arr[0])
	fmt.Println("second item: ", arr[1])

	fmt.Println()

	// len & cap
	fmt.Println("Length of arr: ", len(arr))
	fmt.Println("Capacity of arr: ", cap(arr))

	fmt.Println()

	// Last item
	fmt.Println("Last item of arr: ", arr[len(arr) - 1])

	// change the array items value
	arr[0] = 10
	arr[1] = 20
	arr[len(arr) - 1] = 30
	fmt.Println("arr with new values: ", arr)

	fmt.Println()

	// Ellipsis (...)

	// 2D array

	// slice

	// make

	// create a empty slice

	// Add item to slice

	// append

	// remove item

	// copy in slice

	// sort

	// check a slice is nill (slice == nil)

	fmt.Println()
	fmt.Println("End of array and slice")
}
