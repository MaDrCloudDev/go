package main

import (
	"fmt"
	"sort"
	"strings"
)

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

	// Arrays - fixed length
	var ages [3]int = [3]int{20, 25, 30}

	// shorthand array
	var agesShort = [3]int{21, 26, 31}

	names := [4]string{"Matt", "Mike", "Nicky", "Bob"}
	names[1] = "Michael" // change value

	fmt.Println(ages, len(ages))
	fmt.Println(agesShort, len(agesShort))
	fmt.Println(names, len(names))

	// Slices - like arrays but can be resized
	var scores = []int{100, 50, 60} // no length in brackets
	scores[2] = 25                  // change value
	scores = append(scores, 85)     // add value

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]  // 1st index to 3rd index (not including 3rd)
	rangeTwo := names[2:]   // 2nd index to end
	rangeThree := names[:3] // start to 3rd index (not including 3rd)

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Matt") // add value to rangeOne
	fmt.Println(rangeOne)

	// Standard Library
	// strings package
	greeting := "Hello friends!"

	fmt.Println(strings.Contains(greeting, "Hello"))         // true. strings imported from standard library at top of file
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi")) // doesn't change original string, returns new string, all of these
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll")) // returns index of first instance of "ll"
	fmt.Println(strings.Split(greeting, " "))  // splits string into array of strings

	// sort package (imported at top)
	agesSort := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(agesSort) // sorts array of ints, original slice is changed
	fmt.Println(agesSort)

	index := sort.SearchInts(agesSort, 30) // returns index of value. Searching a value that doesn't exist returns length of array plus 1
	fmt.Println(index)

	namesSort := []string{"Matt", "Mike", "Nicky", "Bob"}
	sort.Strings(namesSort) // sorts in alphabetical order
	fmt.Println(namesSort)

	fmt.Println(sort.SearchStrings(namesSort, "Matt")) // returns index of value

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

// go run main.go
