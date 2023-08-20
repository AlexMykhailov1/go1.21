package main

import (
	"fmt"
	"slices"
)

// Compact removes all duplicate values from a slice.
// Capacity of the previous slice stays the same.
func main() {
	s := []int{1, 1, 2, 2, 3, 4, 5}
	fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
	fmt.Println("s = slices.Compact(s)")
	s = slices.Compact(s)
	fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
}
