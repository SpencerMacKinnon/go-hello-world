package main

import "testing"

// Test Test
func Test(t *testing.T) {
	if t.Name() == "foo" {
		t.Fatal("rip")
	}
}
