package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"log"
	"reflect"
)

func main(){
	fmt.Println("Hello", stuff.Name)
	intArr := []int{1,2,3,4,5}
	strArr := stuff.IntArrToStrArray(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))

	date := stuff.Date{}
	err := date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetYear(1974)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("1st Day : %d/%d/%d\n", date.Day(), date.Month(), date.Year())
}