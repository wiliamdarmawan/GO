package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// var name type 
	// if capital case, it's exported and able to be used by other packages
	// if lower case, only in package scope only

	var vName string = "Wiliam"
	var v1, v2 = 1.2, 3.4
	var v3 = "hello"
	v4 := 5

	pl(vName, v1, v2, v3, v4)
}