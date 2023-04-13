package recursion

func Factorial(num int) int {
	if num <= 1 {
		return 1
	}
	return num * (num - 1)
}
