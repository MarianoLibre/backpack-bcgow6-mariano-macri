package fibotdd


func Fibo(n int) []int {
	a, b := 0, 1
	out := []int{}

	for i := 0; i < n; i++ {
		out = append(out, a)
		temp := a
		a = b
		b = temp + a
	}

	return out
}
