package calculator

func Mul(x int, y int) int {
	return x * y
}

func Div(x int, y int) int {
	if y == 0 {
		panic("Number cannot be devided by 0")
	}

	return x / y
}
