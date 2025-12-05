package main

import "testing"

func TestGetMessage(t *testing.T) {
	expected := "Hello, CI/CD World!"
	actual := GetMessage()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
