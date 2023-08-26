package main

import (
	"fmt"
	"maps"
)

// Clone returns a copy of a given map.
func main() {
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	fmt.Println("m ==", m)
	fmt.Println("maps.Clone(m) ==", maps.Clone(m))
}
