package main

import "fmt"

// Condition Operators (&& , || , > , < , <= , => , !)

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

	fmt.Println()

	// if
	a := 15
	if a > 10 {
		fmt.Println("a is greater than 10")
	}

	fmt.Println()

	// &&
	b := 6
	if b > 0 && b < 10 {
		fmt.Println("b is within range")
	}

	fmt.Println()

	// else
	c := 15
	if c > 20 {
		fmt.Println("c is greater than 20")
	} else {
		fmt.Println("c is less than 20")
	}

	fmt.Println()

	// if, else if, else

	fmt.Println()

	// nested condition

	fmt.Println()

	// statement; condition

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
