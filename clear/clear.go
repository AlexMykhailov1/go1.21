package main

import (
	"fmt"
)

// clear() can be used for a map, slice or a type parameter.
// When used on a map, clear() empties the map.
// For slice, all values are set to zero-values for the corresponding type.
// If the type of the argument is a type parameter,
// all types in its type set must be maps or slices.
func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Println("m ==", m)
	clear(m)
	fmt.Println("clear(m)")
	fmt.Printf("m == nil: %t\n", m == nil)       // false
	fmt.Printf("len(m) == 0: %t\n", len(m) == 0) // true
	fmt.Println("m ==", m)

	s := []int{1, 2}
	fmt.Printf("\ns ==%v\n", s)
	clear(s)
	fmt.Println("clear(s)")
	fmt.Printf("s == nil: %t\n", s == nil)       // false
	fmt.Printf("len(s) == 0: %t\n", len(s) == 0) // false
	fmt.Println("s ==", s)

	//clear(nil) // panic
	fmt.Printf("\nclear(nil) == panic\n")
}
