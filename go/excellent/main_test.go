package main

import "testing"

func TestEvenOrOdd(t *testing.T){
	result := TestEvenOrOdd(12)
	if result != "even" {
		t.Errorf("expected: even, actual: %s", result)
	}
}