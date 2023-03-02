package main

import (
	"fmt"
	"unicode/utf8"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	// In GO, characters are called Runes
	// Runes are unicodes that represent characters
	rStr := "abcdefg"
	pl("Rune Count:", utf8.RuneCountInString(rStr))

	for i, runeVal := range rStr{
		pf("%d : %#U : %c\n", i, runeVal, runeVal) // => %#U => Rune Unicode format
	}
}