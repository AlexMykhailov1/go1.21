package main

import (
	"fmt"
	"slices"
)

// Insert inserts the values v... into s at index i.
// The elements at s[i:] are shifted up.
// In the returned slice r, r[i] == v[0], and r[i+len(v)] == value originally at r[i].
// Function is O(len(s) + len(v)).
func main() {
	s := []int{1, 2, 3}
	fmt.Println("s := []int{1, 2, 3}")
	fmt.Println("s = slices.Insert(s, 0, 4, 5, 6) ==", slices.Insert(s, 0, 4, 5, 6))
	fmt.Println("s = slices.Insert(s, 1, 4, 5, 6) ==", slices.Insert(s, 1, 4, 5, 6))
	fmt.Println("s = slices.Insert(s, 2, 4, 5, 6) ==", slices.Insert(s, 2, 4, 5, 6))
	fmt.Println("s = slices.Insert(s, 3, 4, 5, 6) ==", slices.Insert(s, 3, 4, 5, 6))
	fmt.Println("s = slices.Insert(s, 4, 4, 5, 6) == panic")
}
