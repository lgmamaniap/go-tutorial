package main

import "fmt"

func main() {
	a := 20
	b := 8

	// Sum
	result := a + b
	fmt.Println("Sum: ", result)

	// Subtraction
	result = a - b
	fmt.Println("Subtraction: ", result)

	// Multiplication
	result = a * b
	fmt.Println("Multiplication: ", result)

	// Division
	var div float64 = float64(a) / float64(b)
	fmt.Println("Division: ", div)

	// Modulus
	result = a % b
	fmt.Println("Modulus: ", result)
}
