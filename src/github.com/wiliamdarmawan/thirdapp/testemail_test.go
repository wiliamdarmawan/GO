package app2

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not an email")
	}

	_, err = IsEmail("hello@gmail.com")
	if err != nil {
		t.Error("hello@gmail.com is a valid email")
	}

	_, err = IsEmail("hello@gmail")
	if err == nil {
		t.Error("hello@gmail is not an email")
	}
}