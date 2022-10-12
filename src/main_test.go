// main_test.go

package main

import "testing"

func TestSum(t *testing.T) {
	if add(2, 5) != 7 {
		t.Fail()
	}
	if add(2, 100) != 102 {
		t.Fail()
	}
	if add(222, 100) != 322 {
		t.Fail()
	}
}

func TestProduct(t *testing.T) {
	if multiply(2, 5) != 10 {
		t.Fail()
	}
	if multiply(2, 100) != 200 {
		t.Fail()
	}
	if multiply(222, 3) != 666 {
		t.Fail()
	}
}

func TestSubtract(t *testing.T) {
	if subtract(2, 5) != -3 {
		t.Fail()
	}
	if subtract(100, 100) != 0 {
		t.Fail()
	}
	if subtract(74, 6) != 68 {
		t.Fail()
	}
}
