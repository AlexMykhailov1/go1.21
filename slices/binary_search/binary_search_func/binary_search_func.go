package main

import (
	"cmp"
	"fmt"
	"slices"
)

type tester struct {
	m map[string]int
}

// BinarySearchFunc works like BinarySearch, but uses a custom function for comparison.
// Elements of any types can be compared.
func main() {
	tester1 := tester{m: map[string]int{"Bob": 1, "Alice": 2}}
	tester2 := tester{m: map[string]int{"John": 3, "Paul": 4}}
	tester3 := tester{m: map[string]int{"Martin": 10}}
	unknownTester := tester{m: nil}
	testers := []tester{tester1, tester2, tester3}

	fmt.Println("testers ==", testers)
	fmt.Printf("slices.BinarySearchFunc(testers, tester1, binarySearchCmp) == ") // 0 true
	fmt.Println(slices.BinarySearchFunc(testers, tester1, binarySearchCmp))
	fmt.Printf("slices.BinarySearchFunc(testers, tester3, binarySearchCmp) == ") // 2 true
	fmt.Println(slices.BinarySearchFunc(testers, tester3, binarySearchCmp))
	fmt.Printf("slices.BinarySearchFunc(testers, unknownTester, binarySearchCmp) == ") // 0 false
	fmt.Println(slices.BinarySearchFunc(testers, unknownTester, binarySearchCmp))
}

// Your custom function should return the same values as the cmp.Compare func.
// 0 if the slice element matches the target,
// -1 if the slice element precedes the target,
// 1 if the slice element follows the target.
// Ordering in cmp func should be the same as ordering in the slice.
func binarySearchCmp(t1 tester, t2 tester) int {
	t1Sum := 0
	for _, i := range t1.m {
		t1Sum += i
	}
	t2Sum := 0
	for _, i := range t2.m {
		t2Sum += i
	}
	return cmp.Compare(t1Sum, t2Sum)
}
