package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func getSum2(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	// func funcName(parameters) returnType { BODY }
	pl(getSum2(1, 2, 3, 4, 5))
}