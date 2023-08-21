package main

import (
	"fmt"
	"slices"
)

// IsSorted reports whether x is sorted in ascending order.
func main() {
	fmt.Println("slices.IsSorted([]int{1, 2, 3}) ==", slices.IsSorted([]int{1, 2, 3}))
	fmt.Println("slices.IsSorted([]string{\"C\", \"B\", \"A\"}) ==", slices.IsSorted([]string{"C", "B", "A"}))
}

// You can use IsSortedFunc to understand if your slice is sorted by your own conditions.
