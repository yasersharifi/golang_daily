package main

import "fmt"

func main() {
	fmt.Println("------------------")
	fmt.Println("Switch in golang")
	fmt.Println()

	today := "sat"
	value := isHoliday(today)
	if value {
		fmt.Println("Today is holiday")
	} else {
		fmt.Println("Today is working day")
	}

	fmt.Println()

	// Check variable type name - simple
	var name = "yaser"
	typeNameOfName := typeName(name)
	fmt.Printf("Type of name is: %s", typeNameOfName)
	fmt.Println()

	age := 30
	typeNameOfAge := typeName(age)
	fmt.Printf("Type of age is: %s", typeNameOfAge)
	fmt.Println()

	langs := []string{"Java", "C++", "Go", "C#"}
	typeNameOfLangs := typeName(langs)
	fmt.Printf("Type of langs is: %s", typeNameOfLangs)

	fmt.Println()
	fmt.Println("----------------")
}

// Check day is holiday or not
func isHoliday(day string) bool {
	switch day {
	case "sat", "sun":
		return true
	default:
		return false
	}
}

// Get string of type name
func typeName(variable interface{}) string {
	switch variable.(type) {
	case int:
		return "int"
	case string:
		return "string"
	default:
		return "unknown"
	}
}
