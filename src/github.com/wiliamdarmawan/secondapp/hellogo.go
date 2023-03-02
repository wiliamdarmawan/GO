package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	// var name []datatype
	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Spectacle"
	sl1[4] = "by"
	sl1[5] = "Guy Debord"
	pl("Slice Size:", len(sl1))
	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}

	for _, x := range sl1 {
		pl(x)
	}
	sArr := [5]int { 1, 2, 3, 4, 5 }
	sl2 := sArr[0:2]
	pl("1st 3:", sArr[:3])
	pl("Last 3: ", sArr[2:])
	sArr[0] = 10
	pl("sl2:", sl2)
	sl2[0] = 1
	pl("sArr:", sArr)

	sl2 = append(sl2, 12)
	pl("sl2:", sl2)
	pl("sArr:", sArr)

	sl3 := make([]string, 6)
	pl("sl3:", sl3)
	pl("sl3[0]:", sl3[0])
}