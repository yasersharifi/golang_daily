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

	// by statement
	switch webFramework := "React JS"; webFramework {
	case "React JS", "Next JS":
		fmt.Println("First need to learn React JS.")
	case "Vue JS", "Nuxt Js":
		fmt.Println("First need to learn Vue JS.")
	default:
		fmt.Println("I don't know this framework.")
	}

	fmt.Println()

	// fallthrough
	dayOfWeek := 5

	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
		fallthrough
	case 2:
		fmt.Println("Tuesday")
		fallthrough
	case 3:
		fmt.Println("Wednesday")
		fallthrough
	case 4:
		fmt.Println("Thursday")
		fallthrough
	case 5:
		fmt.Println("Friday")
		fallthrough
	case 6:
		fmt.Println("Saturday")
		fallthrough
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day")
	}

	fmt.Println()

	// More example
	notificationType := "sms" // sms, email, push

	switch notificationType {
	case "sms":
		fmt.Println("Sent sms")
	case "email":
		fmt.Println("Sent email")
	case "push":
		fmt.Println("Sent push")
	default:
		fmt.Println("Choose a correct notification type.")
	}

	fmt.Println()

	// Boolean

	// User status (active: true, false)
	switch isActive := true; isActive {
	case true:
		fmt.Println("User is active.")
	case false:
		fmt.Println("User in't active.")
	default:
		fmt.Println("Unknown")
	}

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
