package main

import "fmt"

func main() {

	studentGrades := map[string]int{
		"Alice": 90,
		"Jane":  85,
		"Dan":   60,
	}
	fmt.Println(studentGrades)

	studentGrades["Alice"] = 95
	fmt.Println(studentGrades)

	alice, ok := studentGrades["Alice"]
	if ok {
		fmt.Println("Alice:", alice)
	}

	key := "Dan"
	if value, ok := studentGrades[key]; ok {
		fmt.Println(key, "=", value)
	}

	delete(studentGrades, key)
	fmt.Println(studentGrades)

	configs := make(map[string]int)
	fmt.Printf("%v %T\n", configs, configs)

	if configs == nil {
		fmt.Println("configs is nil")
	}
}
