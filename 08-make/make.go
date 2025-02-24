package main

import "fmt"

func main() {
	numbers := make([]int, 0, 3)

	// numbers[0] = 100
	// numbers[1] = 200
	// numbers[2] = 300

	numbers = append(numbers, 400)

	fmt.Println(numbers, len(numbers), cap(numbers))
}
