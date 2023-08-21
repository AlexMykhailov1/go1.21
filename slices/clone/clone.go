package main

import (
	"fmt"
	"slices"
)

// Clone returns a copy of the slice.
// The elements are copied using assignment.
func main() {
	s1 := []int{1, 2}
	s2 := []int{3, 4, 5, 6}

	fmt.Printf("s1 := []int{1, 2}\ns2 := []int{3, 4, 5, 6}\n")
	fmt.Println("s1 = slices.Clone(s2)")
	s1 = slices.Clone(s2)
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))
	fmt.Println("assign: s2[3] = 10")
	s2[3] = 10
	fmt.Println("s1[3] still ==", s1[3])
}
