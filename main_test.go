package main

import "testing"

func TestAddition(t *testing.T) {
	a, b := 1, 3
	result := Add(a, b)
	if result < 0 {
		t.Error("Value is negative")
	}
	if result != 4 {
		t.Errorf("%d + %d = %d, Value is not 3", a, b, result)
	} else {
		fmt.Printf("Result : %d is correct!", result)

}
