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

func tbsToML(tbs Tbs) ML {
	return ML(tbs * 14.79)
}

func main() {
	 ml1 := ML(Tsp(3) * 4.92)
	 pf("3 TSPS = %.2f ml\n", ml1)

	 ml2 := ML(Tbs(3) * 14.79)
	 pf("3 Tbs = %.2f ml\n", ml2)

	 pl("2 tsp + 4 tsp = ", tspToML(2) + tspToML(4), "ml")

	 pf("3 Tsp = %.2f ml", tspToML(3))
	 pf("3 Tbs = %.2f ml", tbsToML(3))

}