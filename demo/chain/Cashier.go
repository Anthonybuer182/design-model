package chain

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {
	fmt.Println("结账走人")
}
func (c *Cashier) Next(next Department) {
	c.next = next
}
