package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

type customer struct {
	name string
	address string
	bal float64
}

func getCustInfo(c customer) {
	pf("%s owes us %.2f\n", c.name, c.bal)
}

func newCustAdd(c *customer, address string) {
	c.address = address
}

func main() {
	 var tS customer
	 tS.name = "Wiliam Darmawan"
	 tS.address = "Jl. Kebon Jeruk Raya No. 1"
	 tS.bal = 234.56
	 getCustInfo(tS)
	 newCustAdd(&tS, "123 South st")
	 pl("Address:", tS.address)
	 sS := customer{"Sally Smith", "123 South st", 0.0}

	 pl("Name:", sS.name)
}