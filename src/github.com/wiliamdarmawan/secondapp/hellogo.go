package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

type Tsp float64
type Tbs float64
type ML float64

func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func (tbs Tsp) ToMLs() ML {
	return ML(tbs * 4.92)
}


func (tbs Tbs) ToMLs() ML{
	return ML(tbs * 14.79)
}

func main() {
	tsp1 := Tsp(3)
	pf("%.2f tsp = %.2f ML\n" , tsp1, tsp1.ToMLs())
}