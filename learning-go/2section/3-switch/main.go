package main

import (
	"fmt"
	"time"
)

func main() {

	day := "Sunday"
	fmt.Printf("Today is %s\n", day)

	switch day {
	case "Sunday", "Saturday":
		fmt.Println("It's the weekend!")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday.")
	default:
		fmt.Println("Invalid day.")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("The value is an integer: %d\n", v)
		case string:
			fmt.Printf("The value is a string: %s\n", v)
		case bool:
			fmt.Printf("The value is a boolean: %t\n", v)
		default:
			fmt.Printf("The value is of an unknown type: %T\n", v)
		}
	}

	checkType(42)
	checkType("Hello, Go!")
	checkType(true)
	checkType(3.14)
}
