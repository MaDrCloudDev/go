// Multiple return values
package main

import (
	"strings"
)

func goMultipleReturns(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ") // returns a slice with the first and last name
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 { // check if there are two initials
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

// func main() {
// 	fn1, ln1 := getInitials("Matt Smith")
// 	fmt.Println(fn1, ln1)

// 	fn2, ln2 := getInitials("Nicky Smith")
// 	fmt.Println(fn2, ln2)

// 	fn3, ln3 := getInitials("Bob")
// 	fmt.Println(fn3, ln3)
// }

// go run multipleReturns.go
