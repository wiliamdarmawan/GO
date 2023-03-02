package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	seedSecs := time.Now().Unix()
	randNum := rand.Seed(seedSecs)
	pl("Random:", randNum)
}