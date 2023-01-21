package main

import "testing"

func TestApp(t *testing.T) {
	length := getLength(4)
	if length != 3 {
		t.Errorf("Incorrect")
	}
}
