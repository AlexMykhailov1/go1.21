package main

import (
	"fmt"
	"slices"
)

// Clip removes unused capacity from the slice, returning s[:len(s):len(s)].
func main() {
	s1 := make([]int, 2, 5)
	fmt.Printf("len(s1): %d\ncap(s1): %d\n", len(s1), cap(s1))
	fmt.Println("slices.Clip(s1)")
	slices.Clip(s1) // s1 won't change
	fmt.Printf("len(s1): %d\ncap(s1): %d\n", len(s1), cap(s1))
	fmt.Println("s1 = slices.Clip(s1)")
	s1 = slices.Clip(s1) // s1 is changed
	fmt.Printf("len(s1): %d\ncap(s1): %d\n\n", len(s1), cap(s1))
	s2 := make([]int, 3, 5)
	fmt.Println("s2 := make([]int, 3, 5)")
	fmt.Println("s1 = slices.Clip(s2)")
	s1 = slices.Clip(s2)
	fmt.Printf("len(s1): %d\ncap(s1): %d\n", len(s1), cap(s1))
}
