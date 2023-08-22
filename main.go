package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// strings
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

	// integers
	var ageOne int = 30
	var ageTwo = 40
	ageThree := 50

	fmt.Println(ageOne, ageTwo, ageThree)

	// memory/bits
	var numOne int8 = 127
	var numTwo int16 = 32767
	var numThree int32 = 2147483647
	var numFour int64 = 9223372036854775807
	var numFive uint8 = 255 // unsigned int, can go higher because no negative values

	fmt.Println(numOne, numTwo, numThree, numFour, numFive)

	// floats
	var scoreOne float32 = 12.34
	var scoreTwo float64 = 123456789.123456789
	scoreThree := 1.2 // inferred as float64

	fmt.Println(scoreOne, scoreTwo, scoreThree)

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

// go run main.go
