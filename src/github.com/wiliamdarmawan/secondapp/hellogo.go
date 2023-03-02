package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func getQuotient(x float64, y float64) (ans float64, err error){
	if y == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	} else {
		return x/y, nil
	}
}

func main() {
	// func funcName(parameters) returnType { BODY }
	quotient, err := getQuotient(10,2)
	if err != nil {
		pl(err)
	} else {
		pl(quotient)
	}
}