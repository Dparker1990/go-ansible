package main

import (
	"testing"
)

func TestAcquireUsername(t *testing.T) {
	user := User{}

	user.SetUsername("foo")
	if user.username != "foo" {
		t.Errorf("Username was not set properly")
	}
}

func TestUsernameLengthValidation(t *testing.T) {
	user := User{}
	username := "aabcdefghijklmnopqssrstuvwxzbcdefghijklmnopqrstuvwxz"

	if err := user.SetUsername(username); err == nil {
		t.Errorf("Username must be less than 50 characters")
	}

	if user.username != "" {
		t.Errorf("Username should not have been assigned")
	}
}
