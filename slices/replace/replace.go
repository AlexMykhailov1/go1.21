package main

import (
	"fmt"
	"slices"
)

// Replace replaces the elements s[i:j] by the given v.
func main() {
	s := []int{1, 2, 3}
	fmt.Println("s := []int{1, 2, 3}")
	fmt.Println("slices.Replace(s, 3, 3, 4) ==", slices.Replace(s, 3, 3, 4))
	fmt.Println("slices.Replace(s, 2, 3, 4) ==", slices.Replace(s, 2, 3, 4))
	fmt.Println("slices.Replace(s, 2, 3, 4) ==", slices.Replace(s, 1, 3, 4))
	fmt.Println("slices.Replace(s, 0, 3, 4) ==", slices.Replace(s, 0, 3, 4))

	s2 := []int{5, 6, 7}
	fmt.Printf("\ns2 := []int{5, 6, 7}\n")
	fmt.Println("slices.Replace(s2, 0, 0, 4) ==", slices.Replace(s2, 0, 0, 4))
	fmt.Println("slices.Replace(s2, 0, 1, 4) ==", slices.Replace(s2, 0, 1, 4))
	fmt.Println("slices.Replace(s2, 0, 2, 4) ==", slices.Replace(s2, 0, 2, 4))
	fmt.Println("slices.Replace(s2, 0, 3, 4) ==", slices.Replace(s2, 0, 3, 4))
}
