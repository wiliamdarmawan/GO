package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	pl("Factorial 4 =", factorial(4))
}