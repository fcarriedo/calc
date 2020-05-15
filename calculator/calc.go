package calculator

import "errors"

func Sum(x int, y int) int {
	return x + y
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
