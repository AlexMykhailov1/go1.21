package main

import (
	"cmp"
	"fmt"
)

// Ordered is the type that can be used inside Less, Compare, min and max functions.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func main() {
	// Compare compares 2 ordered values and returns an integer depending on the Comparison.
	fmt.Printf("cmp.Compare(1.5, 2.) == %d\n", cmp.Compare(1.5, 2.)) // -1, x less than y
	fmt.Printf("cmp.Compare(1, 1) == %d\n", cmp.Compare(1, 1))       // 0, x == y
	fmt.Printf("cmp.Compare(2, -2) == %d\n", cmp.Compare(2, -2))     // 1, x greater than y

	// Less tells whether x is less than y.
	fmt.Printf("\ncmp.Less(-1, 0) == %t\n", cmp.Less(-1, 0)) // true
	fmt.Printf("cmp.Less(1., 1.) == %t\n", cmp.Less(1., 1.)) // false
	fmt.Printf("cmp.Less(1, -2) == %t\n", cmp.Less(1, -1))   // false
}
