package main

import (
	"fmt"
	"maps"
)

// Copy copies all key-value pairs in src adding them to dst.
// When a key in src is already present in dst,
// the value in dst will be overwritten by the value associated with the key in src.
func main() {
	m := map[string]int{
		"Mark":   1,
		"Martin": 2,
		"John":   10,
	}
	m1 := map[string]int{
		"John":     3,
		"Chandler": 4,
	}
	fmt.Printf("m := %v\nm1 := %v\n", m, m1)
	fmt.Println("maps.Copy(m, m1)")
	maps.Copy(m, m1)
	fmt.Println("m ==", m)
}
