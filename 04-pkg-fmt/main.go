package main

import "fmt"

func main() {
	hello := "Hello"
	world := " World"
	fmt.Print(hello)
	fmt.Println(world)

	name := "Gus"
	age := 26

	fmt.Printf("My name is %s and I'm %d years old\n", name, age)

	message := fmt.Sprintf("My name is %s and I'm %d years old", name, age)
	fmt.Println(message)

	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of age: %T\n", age)

	var input string
	fmt.Print("Insert a name: ")
	fmt.Scanln(&input)

	fmt.Println("Other name: ", input)
}
