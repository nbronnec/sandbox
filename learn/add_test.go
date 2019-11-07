package main

import (
	"testing"
)

func TestAdd(t *testing.T) {

	t.Log("Add testing\n")

	if 2 != add(1, 1) {
		t.Error("Test failed")
	}

}

func TestOrder(t *testing.T) {
	var low, high int = order(1, 2)
	if low > high {
		t.Errorf("sort error : %v < %v", low, high)
	}
}
