package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Error("The result must be 5")
	}
}

func TestSub(t *testing.T) {
	result := sub(2, 2)

	if result != 0 {
		t.Error("The result must be 0")
	}
}

func TestTimes(t *testing.T) {
	result := times(2, 2)

	if result != 4 {
		t.Error("The result must be 4")
	}
}

func TestSumX(t *testing.T) {
	result := sumX(2, 2)

	if result != 6 {
		t.Error("The result must be 6")
	}
}
