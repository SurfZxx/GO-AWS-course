package main

import "fmt"

func simpleDefer() {
	fmt.Println("Function simpleDefer: Start")
	defer fmt.Println("Function simpleDefer: End")

	fmt.Println("Function simpleDefer: Middle")

}

func lifoSimpleDefer() {
	fmt.Println("Function lastInFirstOut: Start")
	defer fmt.Println("Function lastInFirstOut: First")
	defer fmt.Println("Function lastInFirstOut: Second")
	fmt.Println("Function lastInFirstOut: Middle")

}

func main() {
	//simpleDefer()
	defer func() {
		fmt.Println("Before the return of main()")
	}()

	lifoSimpleDefer()

	fmt.Println("Last In main()")
}
