package main

import (
	"fmt"
)

func goLoops() {
	// loops
	// "while" loop equivalent (but uses "for" in go)
	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	// "for" loop
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	} // does the same thing using for instead of while

	// "for" loop over a slice
	namesLoop := []string{"Matt", "Mike", "Nicky", "Bob"}

	for i := 0; i < len(namesLoop); i++ {
		fmt.Println(namesLoop[i])
	}

	// for-in loop using range
	for index, value := range namesLoop {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	// what if we just want to use the value, not the index? use underscore for unwanted values.
	for _, value := range namesLoop {
		// value = "new string" // doesn't change original array, just this copy of the value(s)
		fmt.Printf("the value is %v \n", value)
	}

	fmt.Println(namesLoop)
}
