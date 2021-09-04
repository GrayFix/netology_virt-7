package main

import (
	"testing"
)

func TestMinFromArray(t *testing.T) {
	val := []int{8, 5, 10}
	min, min_pos := MinFromArray(val)
	if min != 5 && min_pos != 1 {
		t.Error("error in func MinFromArray()")
	}

}
