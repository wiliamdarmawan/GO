package app2

import (
	"fmt"
	"regexp"
)

func IsEmail(s string) (string, error) {
	r, _ := regexp.Compile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]`)

	if r.MatchString(s) {
		return "Valid email", nil
	} else {
		return "", fmt.Errorf("Invalid email")
	}
}