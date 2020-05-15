package calculator

import (
	"testing"
)

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
	if Div(4, 4) != 1 {
		t.Error("Expected 4 / 4 to equal 1")
	}

	if Div(4, 2) != 2 {
		t.Error("Expected 4 / 2 to equal 2")
	}

	if Div(3, 2) != 1 {
		t.Error("Expected 3 / 2 to equal 1")
	}
}
