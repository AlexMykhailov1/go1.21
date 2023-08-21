package main

import (
	"fmt"
	"slices"
)

// Compare goes through a slice starting from index 0,
// and compares elements sequentially using cmp.Compare,
// until non-comparing elements are found.
// Returns the result of cmp.Compare on those elements.
func main() {
	fmt.Println("slices.Compare([]int{1, 2, 3}, []int{1, 2, -3}) ==", slices.Compare([]int{1, 2, 3}, []int{1, 2, -3}))
	fmt.Println("slices.Compare([]int{1, 2, -3}, []int{1, 2, 3}) ==", slices.Compare([]int{1, 2, -3}, []int{1, 2, 3}))
	fmt.Println("slices.Compare([]int{1, 2, 3}, []int{1, 2, 3}) ==", slices.Compare([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println("slices.Compare([]int{1, 2}, []int{1, 2, 3}) ==", slices.Compare([]int{1, 2}, []int{1, 2, 3}))
	fmt.Println("slices.Compare([]int{1, 2, 3}, []int{1, 2}) ==", slices.Compare([]int{1, 2, 3}, []int{1, 2}))
	fmt.Println("slices.Compare([]int{1, 2}, []int{1, 2}) ==", slices.Compare([]int{1, 2}, []int{1, 2}))
}

// Additionally, you can use slices.CompareFunc to compare values with your own function.
