package main

import (
	"testing"
)

func TestMin(t *testing.T) {
	result := Min(2, 3)
	expected := 2
	if result != expected {
		t.Errorf("test Min() fail")
	}
}
