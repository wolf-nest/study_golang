package ch03_test

import "testing"

func TestDataTypeTry(t *testing.T) {
	var a string
	t.Log(a)
	if a == "" {
		t.Log(" ")
	}
}
