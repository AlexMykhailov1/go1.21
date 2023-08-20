package main

import (
	"fmt"
	"slices"
)

// Contains reports whether a value is present in a slice.
// Has O(n) complexity.
func main() {
	fmt.Println("slices.Contains([]int{1}, 1) ==", slices.Contains([]int{1}, 1)) // true
	fmt.Println("slices.Contains([]int{1}, 0) ==", slices.Contains([]int{1}, 0)) // false
	notSoEmpty := make([]int, 1)
	fmt.Println("notSoEmpty := make([]int, 1)")
	fmt.Println("slices.Contains(notSoEmpty, 0) ==", slices.Contains(notSoEmpty, 0)) // true
}
