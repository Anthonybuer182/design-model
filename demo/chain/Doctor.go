package chain

import "fmt"

type Doctor struct {
	next Department
}

func (d *Doctor) Execute(p *Patient) {
	fmt.Println("医生 接待病人")
	d.next.Execute(p)
}
func (d *Doctor) Next(next Department) {
	d.next = next
}
