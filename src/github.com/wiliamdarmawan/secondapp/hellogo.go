package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	// for initialization; condition;
	// postStatement { BODY }

	for x := 1; x <=5; x++{
		pl(x)
	}

	fX := 0
	for fX <= 5 {
		pl(fX)
		fX++
	}
}