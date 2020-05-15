package calculator

import (
	"testing"
)

func TestSum(t *testing.T) {
	if Sum(2, 4) != 6 {
		t.Error("Expected 2 + 4 to equal 6")
	}
}

func TestMul(t *testing.T) {
	if Mul(2, 4) != 8 {
		t.Error("Expected 2 * 4 to equal 8")
	}

	if Mul(2, 4) != 8 {
		t.Error("Expected 2 * 4 to equal 8")
	}

	if Mul(2, 4) != 8 {
		t.Error("Expected 2 * 4 to equal 8")
	}
}

func TestDiv(t *testing.T) {
	if res, _ := Div(4, 4); res != 1 {
		t.Error("Expected 4 / 4 to equal 1")
	}

	if res, _ := Div(4, 2); res != 2 {
		t.Error("Expected 4 / 2 to equal 2")
	}

	if res, _ := Div(3, 2); res != 1 {
		t.Error("Expected 3 / 2 to equal 1")
	}
}

func TestDivByZero(t *testing.T) {
	if _, err := Div(4, 0); err == nil {
		t.Error("Expected error when dividing by 0")
	}
}
