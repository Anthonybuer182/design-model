package strtegy

type Div struct {
}

func (a Div) excute(num1 float64, num2 float64) float64 {
	return num1 / num2
}
