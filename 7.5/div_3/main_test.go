package main

import "testing"

func TestCheckDiv3(t *testing.T) {
	val := 2
	if CheckDiv3(val) {
		t.Error(val, "not divisible by 3 without remainder")
	}
}
