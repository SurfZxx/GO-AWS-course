package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("--- Advanced Slicing Operations ---")
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v, len: %d, cap: %d\n", original, len(original), cap(original))

	s1 := original[2:5] // [2, 3, 4]
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))

	s2 := original[:4] // [0, 1, 2, 3]
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))

	s3 := original[5:] // [5, 6, 7, 8, 9]
	fmt.Printf("s3: %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))

	s4 := original[:] // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	fmt.Printf("s4: %v, len: %d, cap: %d\n", s4, len(s4), cap(s4))

	contains := slices.Contains(s4, 8)
	fmt.Printf("Contains 8: %t\n", contains)
	s4 = slices.Insert(s4, 1, 100)
	fmt.Printf("After inserting 100 at index 1: %v\n", s4)
}
