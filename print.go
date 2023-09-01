package main

import "fmt"

// variables
var ageOne int = 20
var nameOne string = "Matt"
var scoreOne float32 = 25.98

func goPrint() {
	// Print
	fmt.Print("Hello, ")
	fmt.Print("world!") // print on same line

	fmt.Print("Hello, \n")

	// Print line
	fmt.Println("world! \n") // print on new line

	fmt.Println("Hello, world!") // prints on new line automatically
	fmt.Println("New, line!")

	fmt.Println("my age is", ageOne, "and my name is", nameOne) // prints variables

	// Printf (formatted strings) - %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", ageOne, nameOne) // %v = value, values in order

	fmt.Printf("my age is %v and my name is %q \n", ageOne, nameOne)

	fmt.Printf("age is of type %T \n", ageOne) // %T = type

	fmt.Printf("you scored %0.1f points! \n", scoreOne) // %f = float

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", ageOne, nameOne)
	fmt.Println("the saved string is:", str)
}
