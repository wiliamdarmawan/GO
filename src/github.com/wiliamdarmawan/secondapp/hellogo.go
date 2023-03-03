package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

type MyConstraint interface {
	int | float64
}

// pkg.go.dev/golang.org/x/exp/constraints

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {
	pl("5 + 4 =", getSumGen(5, 4))
	pl("5.6 + 4.7 =", getSumGen(5.6, 4.7))
}