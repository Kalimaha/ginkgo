package calculator

import "errors"

func Sum(a int, b int) int {
	return a + b
}

func Divide(a float32, b float32) (err error, result float32) {
	if b == 0 {
		return errors.New("division by 0"), result
	} else {
		return err, a / b
	}
}
