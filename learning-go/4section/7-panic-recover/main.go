package main

import "fmt"

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("Something went wrong in mightPanic")
	}

	fmt.Println("mightPanic did not panic")
}

func recoverable() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from a panic:", err)
		}
	}()

	mightPanic(true)
}

func main() {

	recoverable()

}
