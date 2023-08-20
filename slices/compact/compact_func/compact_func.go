package main

import (
	"fmt"
	"slices"
)

// CompactFunc allows you to use slices.Compact with your own function.
func main() {
	s := []int{-1, -1, -2, -3, -3, 0, 1, 2, 2}
	fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
	fmt.Println("s = removeAllDuplicateNegativeNumbers(s)")
	s = removeAllDuplicateNegativeNumbers(s)
	fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
}

func removeAllDuplicateNegativeNumbers(s []int) []int {
	return slices.CompactFunc(s, func(i int, i2 int) bool {
		if i < 0 && i2 < 0 {
			return i == i2
		}
		return false
	})
}
