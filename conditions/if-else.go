package main

import "fmt"

func main() {
	fmt.Println("-------------")
	fmt.Println("Conditions in golang (if, else, else if)")
	fmt.Println()

	num := 100

	if num > 0 {
		fmt.Println("Number is positive.")
	} else if num < 0 {
		fmt.Println("Number is negative.")
	} else {
		fmt.Println("Number is zero.")
	}

	fmt.Println("----------------")
	fmt.Println("Check number is even or odd")
	fmt.Println()

	num1 := 32
	isEven(num1)
	fmt.Println()

	num2 := 41
	isOdd(num2)
	fmt.Println()

	num3 := 0
	isEven(num3)
	fmt.Println()
	isOdd(num3)

}

func isEven(num int) {
	if num%2 == 0 {
		fmt.Printf("%d is even", num)
		return
	}
	fmt.Printf("%d is odd", num)
}

func isOdd(num int) {
	if num%2 != 0 {
		fmt.Printf("%d is odd", num)
		return
	}
	fmt.Printf("%d is even", num)
}
