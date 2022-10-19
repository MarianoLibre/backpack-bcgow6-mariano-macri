package calculadora

import "errors"


func Restar(num1, num2 int) int {
	return num1 - num2
}

func Dividir(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("Are you f***in' crazy!?")
	}

	return num1 / num2, nil
}
