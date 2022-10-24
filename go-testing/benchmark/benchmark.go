package benchmark


func FactorialIter(n int) int {
	result := 1
	for i := n; i > 1; i-- {
		result *= i
	}

	return result
}


func FactorialRecur(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * FactorialRecur(n - 1)
}
