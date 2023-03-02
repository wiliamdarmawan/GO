package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	// %d : Integer
	// %f : Float
	// %s : String
	// %t : Boolean
	// %v : Value
	// %T : Type
	// %c : Character
	// %o : Base 8
	// %x : Base 16

	pf("%s %d %c %f %t %o %x\n",
		"Hello", 10, 'A', 3.14, true, 10, 10)

	pf("%9f\n", 3.14)
	pf("%.2f\n", 3.1412313)
	pf("%9.f\n", 3.1441231)

	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	pl(sp1)
}