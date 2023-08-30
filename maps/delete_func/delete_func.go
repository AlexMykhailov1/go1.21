package main

import (
	"fmt"
	"maps"
)

// DeleteFunc deletes all key-value pairs that satisfy the given function.
func main() {
	m := map[string]int{
		"1": -1,
		"2": 2,
		"3": 3,
		"4": -4,
	}
	fmt.Println("m :=", m)
	fmt.Println("maps.DeleteFunc(m, deleteAllNegativeValues)")
	maps.DeleteFunc(m, deleteAllNegativeValues)
	fmt.Println("m ==", m)
}

func deleteAllNegativeValues(_ string, v int) bool {
	return v < 0
}
