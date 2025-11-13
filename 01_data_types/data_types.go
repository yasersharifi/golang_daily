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
	fmt.Println()
	fmt.Println("----------- Boolean ----------")
	fmt.Println("Is true: ", isTrue)
	fmt.Println("Is false: ", isFalse)

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
		_  = iota             // Discard the first value (0)
		KB = 1 << (10 * iota) // 1 << (10 * 1)
		MB = 1 << (10 * iota) // 1 << (10 * 2)
		GB = 1 << (10 * iota) // 1 << (10 * 3)
		TB = 1 << (10 * iota) // 1 << (10 * 4)
	)
	fmt.Println()
	fmt.Println("----------- IOTA constant ----------")
	fmt.Println("Sunday: ", SUN) // Output: 0
	fmt.Println("Monday: ", MON) // Output: 1
	fmt.Println("Friday: ", FRI) // Output: 5
	fmt.Println("KB: ", KB)      // Output: 1024
	fmt.Println("MB: ", MB)      // Output: 1048576
	fmt.Println("GB: ", GB)
	fmt.Println("TB: ", TB)

	// nil
	var arr []int = nil
	fmt.Println()
	fmt.Println("----------- Nil ----------")
	fmt.Println("Nil Array: ", arr)
	fmt.Println("Is Nil: ", arr == nil)

	/* Types (تایپ ها)
	int, int8, int16, int32, int64, uint,
	uint8, uint16, uint32, uint64, uintptr,
	float32, float64, complex128, complex64,
	bool, byte, rune, string, error
	*/
	var x int = 23                     // int
	var x1 int8 = 100                  // int8 (-128, 127)
	var x2 int16 = 1000                // int16 (-32768, 32767)
	var x3 int32 = 2147483640          // int16 (-2147483648, 2147483647)
	var x4 int64 = 9223372036854775800 // int16 (-9223372036854775808, 9223372036854775807)
	fmt.Println()
	fmt.Println("----------- Integer ----------")
	fmt.Println("x(int): ", x)
	fmt.Println("x1(int8): ", x1)
	fmt.Println("x2(int16): ", x2)
	fmt.Println("x3(int32): ", x3)
	fmt.Println("x4(int64): ", x4)

	var y uint = 20
	var y1 uint8 = 225                  // unsigned int (0, 255)
	var y2 uint16 = 6553                // unsigned int (0, 65535)
	var y3 uint32 = 429496729           // unsigned int (0, 4294967295)
	var y4 uint64 = 1844674407370955161 // unsigned int (0, 18446744073709551615)
	fmt.Println()
	fmt.Println("----------- Unsigned Integer ----------")
	fmt.Println("y(uint): ", y)
	fmt.Println("y1(uint8): ", y1)
	fmt.Println("y2(uint16): ", y2)
	fmt.Println("y3(uint32): ", y3)
	fmt.Println("y4(uint64): ", y4)

	var firstName = "Yaser"
	var lastName = "Sharifi"
	var email = "yassersharifi@74gmail.com"
	fmt.Println()
	fmt.Println("----------- String ----------")
	fmt.Println("First Name is: ", firstName)
	fmt.Println("Last Name is: ", lastName)
	fmt.Println("Full Name is: ", firstName+" "+lastName)
	fmt.Println("Email is: ", email)

	/* Functions (توابع)
	make, len, cap, new, append, copy, close,
	delete, complex, real, imag, panic, recover
	*/

	fmt.Println()
	fmt.Println("----------------------")
	fmt.Println("End of data types.")
}
