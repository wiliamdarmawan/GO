package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

type contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name string
	address string
	contact
}

func (b business) info() {
	pf("Contact at %s is %s %s", b.name, 
	b.contact.fName, b.contact.lName)
}

func main() {
	con1 := contact { 
		fName: "Wiliam", 
		lName: "Darmawan", 
		phone: "1234567890",
	}
	biz1 := business { 
		name: "Wiliam's Business", 
		address: "123 Main St", 
		contact: con1,
	}
	biz1.info()
}