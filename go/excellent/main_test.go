package main

import "testing"

func TestEvenOrOdd(t *testing.T){
	result := EventOrOdd(10)
	if result != "even"{
		t.Errorf("expected: even, actual: %s",result)
	}
}