package main

import "fmt"

func goArrays() {
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
}
