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
	fmt.Println("Last item of arr: ", arr[len(arr)-1])

	// change the array items value
	arr[0] = 10
	arr[1] = 20
	arr[len(arr)-1] = 30
	fmt.Println("arr with new values: ", arr)

	fmt.Println()

	// Ellipsis (...)

	// 2D array

	// slice
	slice := []string{"A", "B", "C"}
	fmt.Println("Slice: ", slice)
	fmt.Println("First item of Slice: ", slice[0])

	// append to slice
	slice = append(slice, "D")
	slice = append(slice, "E")
	fmt.Println("Slice with new items: ", slice)

	fmt.Println()

	// make
	langs := make([]string, 3)
	langs[0] = "C++"
	langs[1] = "Go"
	fmt.Println("Langs (by make): ", langs)
	fmt.Println("Length of langs (by make): ", len(langs))
	fmt.Println("Capacity of langs (by make): ", cap(langs))

	fmt.Println()

	// create a empty slice
	emptySlice := []int{}
	fmt.Println("Empty slice: ", emptySlice)

	// Add item to slice
	emptySlice = append(emptySlice, 10)
	emptySlice = append(emptySlice, 30)
	fmt.Println("New item in slice: ", emptySlice)

	// check a slice is nil (slice == nil)
	var s []int
	s2 := []int{}
	fmt.Println("Check s is nil: ", s == nil)   // true
	fmt.Println("Check s2 is nil: ", s2 == nil) // false

	// remove item

	// copy in slice

	// sort

	fmt.Println()
	fmt.Println("End of array and slice")
}
