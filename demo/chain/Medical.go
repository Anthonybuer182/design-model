package chain

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) Execute(p *Patient) {
	fmt.Println("药房处 接待病人")
	m.next.Execute(p)
}
func (m *Medical) Next(next Department) {
	m.next = next
}
