package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	pl("5 + 4", 5+4)
	pl("5 - 4", 5-4)
	pl("5 * 4", 5*4)
	pl("5 / 4", 5/4)
	pl("5 % 4", 5%4)
	mInt := 1
	mInt += 1
	mInt++
	pl(mInt)
}