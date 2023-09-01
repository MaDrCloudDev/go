package main

import (
	"fmt"
)

func goBooleansAndConditionals() {
	// booleans and conditionals
	booleanAge := 18

	fmt.Println(booleanAge <= 18) // true

	// if statement
	if booleanAge < 18 {
		fmt.Println("You can't vote!")
	} else if booleanAge <= 18 {
		fmt.Println("You can vote!")
	} // else...

	// cycle through slice
	booleanNames := []string{"Matt", "Mike", "Nicky", "Bob"}

	for index, value := range booleanNames {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue // the continue keyword skips the rest of the code in the loop and goes to the next iteration
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break // the break keyword breaks out of the loop
		}

		fmt.Printf("the value at index %v is %v \n", index, value)
	}
}
