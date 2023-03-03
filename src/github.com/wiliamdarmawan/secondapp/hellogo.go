package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	// var myMap map [keyType]valueType
	var heroes map[string]string
	heroes = make(map[string]string)
	villains := make(map[string]string)
	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["Flash"] = "Barry Allen"
	villains["Lex Luthor"] = "Lex Luthor"

	superPets := map[int]string {
		1: "Krypto",
		2: "Bat Hound",
	}

	pf("Batman is %v\n", heroes["Batman"])
	pl("Chip:", superPets[3])

	_, ok := superPets[3]
	 pl("Is there a 3rd pet:", ok)

	 for k, v := range heroes {
		pf("%s is %s\n", k, v)
	 }
	 delete(heroes, "Flash")
	 pl(heroes)
}