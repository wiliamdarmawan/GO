package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func changeVal2(myPtr *int) {
	*myPtr = 12
}

func main() {
	// func funcName(parameters) returnType { BODY }
	f3 := 5
	pl("f3 before func:", f3)
	changeVal2(&f3)
	pl("f3 after func:", f3)
}