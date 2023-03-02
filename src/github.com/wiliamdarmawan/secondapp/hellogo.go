package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func changeVal2(myPtr *int) {
	*myPtr = 12
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var NumSize float64 = float64(len(nums))

	for _, val := range nums {
		sum += val
	}
	return (sum / NumSize)
}

func main() {
	// func funcName(parameters) returnType { BODY }
	iSlice := []float64{11,13,17}
	pf("Average: %.3f\n", getAverage(iSlice...))
}