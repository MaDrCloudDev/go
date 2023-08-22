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
}

// go run main.go
