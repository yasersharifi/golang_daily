package main

import "fmt"



func main() {
	fmt.Println("Data types in golang: ")
	fmt.Println("----------------------")
	fmt.Println()

	/* Constant (ثابت ها)
		true, false, iota, nil
	*/
	isTrue := true
	isFalse := false
	const (
		SUN = iota
		MON
		TUE
		WED
		THU
		FRI
		SAT
	)
	const (
		_ = iota // Discard the first value (0)
		KB = 1 << (10 * iota) // 1 << (10 * 1)
		MB = 1 << (10 * iota) // 1 << (10 * 2)
		GB = 1 << (10 * iota) // 1 << (10 * 3)
		TB = 1 << (10 * iota) // 1 << (10 * 4)
	)
	// nil
	var arr []int = nil

	fmt.Println("Is true: ", isTrue)
	fmt.Println("Is false: ", isFalse)
	fmt.Println("Sunday: ", SUN) // Output: 0
	fmt.Println("Monday: ", MON) // Output: 1
	fmt.Println("Friday: ", FRI) // Output: 5
	fmt.Println("KB: ", KB) // Output: 1024
	fmt.Println("MB: ", MB) // Output: 1048576
	fmt.Println("GB: ", GB)
	fmt.Println("TB: ", TB)
	fmt.Println("Nil Array: ", arr)
	fmt.Println("Is Nil: ", arr == nil)

	/* Types (تایپ ها)
		int, int8, int16, int32, int64, uint,
		uint8, uint16, uint32, uint64, uintptr,
		float32, float64, complex128, complex64,
		bool, byte, rune, string, error
	*/

	/* Functions (توابع)
		make, len, cap, new, append, copy, close,
		delete, complex, real, imag, panic, recover
	*/
}

