package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]

	return
}

func main() {

	fmt.Println(time.Now())

	fmt.Println(divide(4, 0))

	value, err := divide(4, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	firstName, LastName := splitName("Hugo Sa")
	fmt.Println(firstName, LastName)
}
