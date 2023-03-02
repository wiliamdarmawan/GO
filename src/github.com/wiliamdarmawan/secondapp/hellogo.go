package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func getArraySum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	// func funcName(parameters) returnType { BODY }
	vArr := []int{1, 2, 3, 4, 5}
	pl("Array Sum:", getArraySum(vArr))
}