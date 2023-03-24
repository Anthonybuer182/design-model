package chain

type Department interface {
	Next(Department)
	Execute(*Patient)
}
