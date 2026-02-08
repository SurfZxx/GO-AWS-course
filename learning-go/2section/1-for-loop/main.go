package main

import "fmt"

func main() {

	//c style loop
	for i := 0; i < 10; i++ {
		//fmt.Println(i)
	}

	//while style
	k := 3
	for k > 0 {
		//fmt.Println(k)
		k--
	}

	//infinte loop
	fmt.Println("----Infinite loop ----")
	counter := 0
	for {
		fmt.Println(counter)
		counter++
		if counter == 10 {
			break
		}
	}

	fmt.Println("---Skipping numbers---")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("----For Each----")
	items := [3]string{"GO", "Python", "Java"}
	for index, value := range items {
		fmt.Println(index, value)
	}
	//se nao queremos index ou value, omitir usando _ (underscore) no seu lugar
}
