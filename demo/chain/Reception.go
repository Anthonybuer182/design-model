package chain

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Patient) {
	fmt.Println("接待处 接待病人")
	r.next.Execute(p)
}
func (r *Reception) Next(next Department) {
	r.next = next
}
