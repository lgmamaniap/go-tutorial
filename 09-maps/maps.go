package main

import "fmt"

func main() {
	days := make(map[int]string)

	days[0] = "Sunday"
	days[1] = "Monday"
	days[2] = "Tuesday"
	days[3] = "Wednesday"
	days[4] = "Thursday"
	days[5] = "Friday"
	days[6] = "Saturday"
	fmt.Println(days, len(days))

	days[7] = "SATURDAY"
	fmt.Println(days, len(days))

	// Delete element
	delete(days, 7)
	fmt.Println(days, len(days))

	// Students map
	students := make(map[string][]int)

	students["Gus"] = []int{10, 9, 8, 7, 6}
	students["Tavo"] = []int{9, 8, 7, 6, 5}

	fmt.Println(students)
	fmt.Println(students["Gus"])

	fmt.Println(students["Gus"][1])

}
