package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)

	cV3 := "5000000000"
	cV4, err := strconv.Atoi(cV3) // Atoi => ASCII to Integer
	pl(cV4, err, reflect.TypeOf(cV4))

	cV5 := 5000000000
	cV6 := strconv.Itoa(cV5) // Itoa => integer to ascii

	pl(cV6, reflect.TypeOf(cV6))

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8, reflect.TypeOf(cV8))
	} else
	{
		log.Fatal(err)
	}
	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9, reflect.TypeOf(cV9))
}