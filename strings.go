package main

import (
	"fmt"
)

func goStrings() {
	// strings
	fmt.Println("Hello, World!")

	var nameOne string = "Matt"
	fmt.Println(nameOne)

	// type inferred
	var nameTwo = "Mike"
	fmt.Println(nameTwo)

	// preset variables
	var nameThree string
	fmt.Println(nameThree) // empty string

	nameOne = "Bill"
	nameThree = "Bob"

	// shorthand (can't be used outside of a function)
	nameFour := "Bobby"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)
}
