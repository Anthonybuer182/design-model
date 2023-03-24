package adapter

import "fmt"

type MacTypec struct {
}

func (mac MacTypec) insert() {
	fmt.Println("typec的可直接接入")
}
