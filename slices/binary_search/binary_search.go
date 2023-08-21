package main

import (
	"fmt"
	"slices"
)

// BinarySearch searches slice for target which is of cmp.Ordered type.
// Returns a target index and a bool, indicating whether the target was found.
// Slice must be sorted in an increasing order.
// For more examples visit: https://pkg.go.dev/slices@master#BinarySearch.
func main() {
	sorted := []int{1, 2, 3, 4}
	fmt.Printf("sorted == %v\n", sorted)
	fmt.Print("slices.BinarySearch(sorted, 1) == ") // 0 true
	fmt.Println(slices.BinarySearch(sorted, 1))
	fmt.Print("slices.BinarySearch(sorted, 4) == ") // 3 true
	fmt.Println(slices.BinarySearch(sorted, 4))
	fmt.Printf("slices.BinarySearch(sorted, -1) == ") // 0 false
	fmt.Println(slices.BinarySearch(sorted, -1))

	// Try to run search on not sorted slice
	unsorted := []int{4, 3, 2, 1}
	fmt.Printf("\nunsorted == %v\n", unsorted)
	fmt.Print("slices.BinarySearch(unsorted, 1) == ") // 0 false (unexpected)
	fmt.Println(slices.BinarySearch(unsorted, 1))
	fmt.Print("slices.BinarySearch(unsorted, 4) == ") // 4 false (unexpected)
	fmt.Println(slices.BinarySearch(unsorted, 4))
	fmt.Print("slices.BinarySearch(unsorted, -1) == ") // 0 false (unexpected)
	fmt.Println(slices.BinarySearch(unsorted, -1))
}
