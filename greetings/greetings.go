package greetings

import "fmt"

var points = []int{20, 90, 100, 45, 70}

var score = 99.5

// cannot use shorthand outside of functions
// scoreTwo := 50

func sayHello(n string) {
	fmt.Println("hello", n)
}

func showScore() {
	fmt.Println("You scored this many points:", score)
}
