package strtegy

type Add struct {
}

func (a Add) excute(num1 float64, num2 float64) float64 {
	return num1 + num2
}
