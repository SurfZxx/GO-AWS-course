package main

import "fmt"

func modifyValue(val int) {
	val = val * 10
	fmt.Printf("modifyValue: %+v\n", val)
}

func modifyPointer(val *int) {
	if val == nil {
		fmt.Printf("modifyPointer: val is nil\n")
	}
	*val = *val * 10
	fmt.Printf("modifyPointer: %+v\n", *val)
}

func main() {

	age := 10
	//agePtr := &age

	modifyValue(age)
	fmt.Printf("modifyValue: %+v\n", age)

	modifyPointer(&age)
	fmt.Printf("agePtr: %+v\n", age)

	grade := 50
	gradePtr := &grade
	fmt.Printf("gradePtr: %+v\n", gradePtr)
	fmt.Printf("gradePtr: %+v\n", &gradePtr)
}
