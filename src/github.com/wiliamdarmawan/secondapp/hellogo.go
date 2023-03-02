package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	// Conditional Operators: ==, !=, <, >, <=, >=
	// Logical Operators : && || !

	iAge := 8
	if iAge >= 18 {
		pl("You can vote!")
	} else {
		pl("You can't vote!")
	}

	pl("!true =", !true)
}