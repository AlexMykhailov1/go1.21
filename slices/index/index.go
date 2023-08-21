package main

import (
	"fmt"
	"slices"
)

// Index returns the index of the first occurrence of v in s, or -1 if not present.
// Has O(n) complexity.
func main() {
	s := []int{1, 2, 3}
	fmt.Println("s := []int{1, 2, 3}")
	fmt.Println("slices.Index(s, 1) ==", slices.Index(s, 1)) // 0
	fmt.Println("slices.Index(s, 3) ==", slices.Index(s, 3)) // 2
	fmt.Println("slices.Index(s, 4) ==", slices.Index(s, 4)) // -1
}

// You can use an IndexFunc to find an index of the first element
// that satisfies your own conditions.
