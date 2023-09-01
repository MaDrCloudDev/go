package main

import (
	"fmt"
	"sort"
	"strings"
)

func goStdLibrary() {
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
}
