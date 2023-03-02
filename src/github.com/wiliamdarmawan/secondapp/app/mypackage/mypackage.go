package stuff

import (
	"strconv"
)

var Name string = "Wiliam Darmawan"

 func IntArrToStrArray(arr []int) []string {
	var strArr []string
	for _, i := range arr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr
 }