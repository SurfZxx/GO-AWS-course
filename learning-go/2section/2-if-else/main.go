package main

import "fmt"

func main() {

	tmp := 25
	if tmp > 30 {
		fmt.Println("greater than 30")
	} else {
		fmt.Println("lesser than 30")
	}

	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: You suck!")
	}

	userAccess := map[string]bool{
		"jane": true,
		"john": false,
	}

	if hasAccess, ok := userAccess["jane"]; ok && hasAccess {
		fmt.Println("Jane can have access")
	} else {
		fmt.Println("Jane do not have access")
	}
	//duas maneiras de fazer
	hasAccessm, ok := userAccess["jane"]
	if ok && hasAccessm {
		fmt.Println("Jane has access!")
	} else {
		fmt.Println("Jane has no access!")
	}
}
