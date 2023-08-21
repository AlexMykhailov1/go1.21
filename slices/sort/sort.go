package main

import (
	"fmt"
	"slices"
)

// Sort sorts a slice of any ordered type in ascending order.
func main() {
	s := []int{-1, 2, 46, 123, -32, 23, 4, 23}
	fmt.Println("slices.Sort(s)")
	slices.Sort(s)
	fmt.Println("s ==", s)
}

// For Sort function there are also SortFunc and SortStableFunc.

// SortFunc can be used to sort slices with your own cmp function.
// This sort is not guaranteed to be stable.
// Also, SortFunc requires that cmp is a strict weak ordering.
// See https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings.

// SortStableFunc is the same as SortFunc,
// but sorts the slice x while keeping the original order of equal elements.
