package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	sV1 := "A word"
	// strings are just array of bytes ([]bytes)
	
	pl(sV1)
	replacer := strings.NewReplacer("A", "Another", "word", "sentence")
	sV2 := replacer.Replace(sV1)
	pl(sV2)

	pl("Length:", len(sV2))
	pl("Contains Another:", strings.Contains(sV2, "Another"))
	pl("o index:", strings.Index(sV2, "o"))
	pl("Replace:", strings.Replace(sV2, "o", "0", -1))

	sV3 := "\nSome Words\n" // \t \" \\
	sV3 = strings.TrimSpace(sV3)
	pl("Split:", strings.Split("a-b-c-d", "-"))
	pl("Lower:", strings.ToLower(sV3))
	pl("Upper:", strings.ToUpper(sV3))
	pl("Prefix:", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix:", strings.HasSuffix("tacocat", "cat"))
}