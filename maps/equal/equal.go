package main

import (
	"fmt"
	"maps"
)

// Equal tells whether 2 maps contain the same key-value pairs.
// You can also use an EqualFunc to compare maps with your own function.
func main() {
	m := map[string]int{
		"Bob":  1,
		"Mary": 2,
	}
	m1 := map[string]int{
		"Bob":  10,
		"Mary": 2,
	}
	fmt.Println("m :=", m)
	fmt.Println("m1 :=", m1)
	fmt.Println("maps.Equal(m, m1) ==", maps.Equal(m, m1))
}
