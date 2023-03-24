package observer

import "fmt"

type Customer struct {
	Id   int
	Name string
}

func (c Customer) subscribe(msg string) {
	fmt.Printf("%v接受的消息为%v", c.Name, msg)
}
func (c Customer) getId() int {
	return c.Id
}
