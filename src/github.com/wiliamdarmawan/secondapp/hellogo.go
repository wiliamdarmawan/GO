package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr{
		pf("Rune Array: %d\n", v)

	}

	byteArr := []byte{'a', 'b', 'c' }
	bStr := string(byteArr[:])
	pl("I'm  a string: ", bStr)
}