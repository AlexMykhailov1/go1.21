package main

import (
	"fmt"
)

// Min and max functions use the same logic under the hood,
// the only difference being the values they search for.
// min and max functions can only compare types that inherit cmp.Ordered interface,
// see more in the cmp package.
func main() {
	a, b := 1, 2
	c := min(a)              // c == a
	d := max(a, b)           // d == max(a,b)
	e := max(float64(a), 52) // e is of type float64
	//f := min(float64(a), b, 52) // compilation error: b is of type int

	// string operations are compared lexically byte-wise
	g := max("abc", "def")

	fmt.Printf("min(a) == %v\n", c)
	fmt.Printf("Type(min(a)) == %T\n", c)
	fmt.Printf("\nmax(a,b) == %d\n", d)
	fmt.Printf("Type(max(a,b)) == %T\n", d)
	fmt.Printf("\nmax(float64(a), b, 52) == %.1f\n", e)
	fmt.Printf("Type(max(float64(a), b, 52)) == %T\n", e)
	fmt.Printf("\nmax(\"abc\",\"def\") == %s\n", g)
	fmt.Printf("Type(max(\"abc\",\"def\")) == %T\n", g)
}
