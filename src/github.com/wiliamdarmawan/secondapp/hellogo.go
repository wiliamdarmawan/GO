package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func changeVal2(myPtr *int) {
	*myPtr = 12
}

func dblArrVals(arr *[4]int) {
	for i := 0; i < 4; i++ {
		arr[i] *= 2
	}
}

func main() {
	// func funcName(parameters) returnType { BODY }
	pArr := [4]int{1, 2, 3, 4}
	dblArrVals( &pArr ) // double array values
	pl(pArr)
}