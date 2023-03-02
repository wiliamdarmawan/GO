package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50)

	for true {
		pl("Guess the number between 0 and 50: ")
		pl("Random Number is:", randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		guess = strings.TrimSpace((guess))
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)	
		}else if iGuess > randNum {
			pl("Pick a lower value")
		}else if iGuess < randNum{
			pl("Pick a higher value")
		} else{
			pl("You guesssed it")
			break
		}
	}
}