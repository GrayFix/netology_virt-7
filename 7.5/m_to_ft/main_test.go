package main

import "testing"

func TestM_to_ft(t *testing.T) {
	val := 0.3048
	if m_to_ft(val) != 1 {
		t.Error(val, "for some reason, it is not equal to 1 ft")
	}
}
