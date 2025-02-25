package main

import (
	"fmt"
	"slices"
)

func main() {
	// slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers, len(numbers))

	// Add element
	numbers = append(numbers, 6)
	numbers = append(numbers, 7)
	numbers = append(numbers, 8)
	fmt.Println(numbers, len(numbers))

	// Sub slice
	subNumbers := numbers[:2]
	fmt.Println(subNumbers, len(subNumbers))

	numbers[0] = 100

	fmt.Println(numbers, len(numbers))
	fmt.Println(subNumbers, len(subNumbers))

	months := []string{"January", "February", "March"}
	fmt.Printf("Len: %d, Cap: %d, p: %p\n", len(months), cap(months), months)

	months = append(months, "April")
	fmt.Printf("Len: %d, Cap: %d, p: %p\n", len(months), cap(months), months)

	testData := []string{"a", "b", "c", "d", "e"}
	fmt.Println(slices.Contains(testData, "a"))
	fmt.Println(slices.Contains(testData, "z"))
}
