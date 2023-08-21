package main

import (
	"fmt"
	"slices"
)

// Min and Max functions find a minimum and maximum values in a slice.
func main() {
	s := []int{1, 2, 3}
	fmt.Println("s := []int{1, 2, 3}")
	fmt.Println("slices.Max(s) ==", slices.Max(s))
	fmt.Println("slices.Min(s) ==", slices.Min(s))
}

// There are also MinFunc and MaxFunc,
// which can be used to find minimum and maximum values with your own function.
