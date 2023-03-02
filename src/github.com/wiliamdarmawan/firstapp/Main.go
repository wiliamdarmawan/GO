package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i) // itoa => integer to ascii
	fmt.Printf("%v, %T\n", j, j)
}