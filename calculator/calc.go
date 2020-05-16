package calculator

import (
	"errors"
	"github.com/thatisuday/nummanip/calc"
)

// Sums all the numbers given as arguments
func Sum(nums ...int) int {
	return calc.Add(nums...)
}

// Multiplies the 2 numbers given as arguments
func Mul(x int, y int) int {
	return x * y
}

// Divides x / y.
// Will return error if y = 0 (as div by Zero is undefined)
func Div(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Division by 0 is undefined")
	}

	return x / y, nil
}
