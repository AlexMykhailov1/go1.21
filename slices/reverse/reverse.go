package main

import (
	"fmt"
	"slices"
)

// Reverse reverses elements of a slice.
func main() {
	s := []int{1, 2, 3}
	fmt.Println("s := []int{1, 2, 3}")
	fmt.Println("slices.Reverse(s)")
	slices.Reverse(s)
	fmt.Println("s ==", s)

	s2 := []int{5, 6, 7}
	fmt.Printf("\ns2 := []int{5, 6, 7}\n")
	fmt.Println("slices.Reverse(s2[0:2])")
	slices.Reverse(s2[0:2])
	fmt.Println("s2 ==", s2)
}
