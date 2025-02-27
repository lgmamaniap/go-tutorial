package main

import (
	"errors"
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	q := &j        // point to j
	*q = *p + 10   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// sumNums
	result, err := sumNums(*p, *q)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func sumNums(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a and b must be positive numbers")
	}
	return a + b, nil
}
