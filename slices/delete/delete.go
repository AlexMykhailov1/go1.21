package main

import (
	"fmt"
	"slices"
)

// Delete removes the elements s[i:j] from s.
// A panic is received when s[i:j] is not a valid slice.
// Delete has O(len(s)-j) complexity. (Still O(n))
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("numbers ==", numbers)
	fmt.Println("numbers = slices.Delete(numbers, 1, 3)")
	numbers = slices.Delete(numbers, 1, 3)
	fmt.Println("numbers ==", numbers)
}

// You can also use DeleteFunc to delete elements that satisfy your function.
