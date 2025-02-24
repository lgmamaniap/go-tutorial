package main

import "fmt"

func main() {
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println(numbers)
	fmt.Println(numbers[4])

	// Array with numbers
	numbers2 := [3]string{"one", "two", "three"}
	fmt.Println(numbers2)

	// Array without define the length
	colors := [...]string{
		"red",
		"green",
		"blue",
		"yellow",
		"black",
		"white",
	}
	fmt.Println(colors, len(colors))

	// Array of money, define values by index
	money := [...]string{
		0: "Dollar",
		2: "Euro",
		3: "Peso",
		5: "Yen",
	}
	fmt.Println(money, len(money))

	// Sub arrays
	subMoney := money[3:]
	fmt.Println(subMoney)
}
