package main

import (
	"fmt"
	"slices"
)

// Equal is similar to slices.Compare, but instead of a cmp.Compare computed value returns a bool.
// Comparison stops after the first non-equal pair is found.
// Has complexity of O(n).
func main() {
	fmt.Println("slices.Equal([]int{1, 2, 3}, []int{1, 2, -3}) ==", slices.Equal([]int{1, 2, 3}, []int{1, 2, -3}))
	fmt.Println("slices.Equal([]int{1, 2, 3}, []int{1, 2, 3}) ==", slices.Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println("slices.Equal([]int{1, 2}, []int{1, 2, 3}) ==", slices.Equal([]int{1, 2}, []int{1, 2, 3}))
	fmt.Println("slices.Equal([]int{1, 2}, []int{1, 2}) ==", slices.Equal([]int{1, 2}, []int{1, 2}))
}

// Same as with slices.CompareFunc, you can use EqualFunc to compare slice values as you wish.
