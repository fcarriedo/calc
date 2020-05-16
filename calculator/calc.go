package calculator

import (
	"errors"
	"github.com/thatisuday/nummanip/calc"
)

func Sum(nums ...int) int {
	return calc.Add(nums...)
}

func Mul(x int, y int) int {
	return x * y
}

func Div(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Division by 0 is undefined")
	}

	return x / y, nil
}
