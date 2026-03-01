package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

func NewEmployee(id int, firstName string, lastName string, position string, salary int, isActive bool) Employee {
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		Salary:    salary,
		IsActive:  isActive,
		JoinedAt:  time.Now(),
	}
}

func main() {

	jane := Employee{
		ID:        1,
		FirstName: "Jane",
		LastName:  "Doe",
		Position:  "CFO",
		Salary:    1000,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}
	fmt.Println(jane)

	joe := NewEmployee(2, "John", "Doe", "CFO", 20000, true)
	fmt.Println(joe)

	joePtr := &joe
	joePtr.Salary = 15000
	fmt.Println(*joePtr)
}
