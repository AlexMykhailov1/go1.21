package main

import (
	"fmt"
	"slices"
)

// Grow adds a minimum of n capacity to the slice.
// After Grow, at least n values can be added to the slice without allocating.
// Panics if n is negative or is too large to allocate memory.
func main() {
	s := make([]int, 2)
	fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
	fmt.Println("s = slices.Grow(s, 3)")
	s = slices.Grow(s, 3)
	fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
}
