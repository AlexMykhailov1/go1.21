package main

import (
	"fmt"
	"slices"
)

// ContainsFunc lets us use a custom function to find element in a slice.
func main() {
	apple := food{t: foodTypeFruit, name: "apple"}
	pear := food{t: foodTypeFruit, name: "pear"}

	fruits := []food{apple, pear}
	fmt.Println("fruits ==", fruits)
	fmt.Println("containsFruit(fruits) ==", containsFruit(fruits))
}

type food struct {
	t    foodType
	name string
}

type foodType string

const foodTypeFruit foodType = "fruit"

func containsFruit(fd []food) bool {
	return slices.ContainsFunc(fd, func(f food) bool {
		return f.t == foodTypeFruit
	})
}
