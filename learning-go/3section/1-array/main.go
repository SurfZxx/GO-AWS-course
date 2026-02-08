package main

import "fmt"

func main() {

	var numbers [2]int

	fmt.Printf("%+v\n", numbers)

	numbers[0] = 1
	numbers[1] = 2
	fmt.Println(numbers)

	primes := [4]int{2, 3, 5, 7}
	fmt.Println(primes)

	primes[3] = 11
	fmt.Println(primes)

	for i := 0; i < len(primes); i++ {
		fmt.Println(primes[i])
	}

	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[1][2] = 3
	fmt.Println(matrix)

}
