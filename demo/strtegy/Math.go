package strtegy

type Math struct {
}

func Count(c Calculate, num1 float64, num2 float64) float64 {
	return c.excute(num1, num2)
}
