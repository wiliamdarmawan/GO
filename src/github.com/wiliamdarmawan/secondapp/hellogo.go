package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	/*
		Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.

		O_RDONLY : open the file read-only.
		O_WRONLY : open the file write-only.
		O_RDWR : open the file read-write.

		These can be or'ed

		O_APPEND : append data to the file when writing.
		O_CREATE : create a new file if none exists.
		O_EXCL : used with O_CREATE, file must not exist.
		O_SYNC : open for synchronous I/O.
		O_TRUNC : if possible, truncate file when opened.
	*/

	_, err := os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		pl("File doesn't exist")
	} else {
		f, err := os.OpenFile("data.txt", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}

		
	}
}