package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello %s\n", name)
}

func sum(a, b int) int {
	return a + b
}

func main() {
	greet("Gus")
	resp := sum(10, 20)
	fmt.Printf("Sum of two numbers: %d\n", resp)
}
