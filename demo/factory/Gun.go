package factory

import "fmt"

type Gun struct {
	Name  string
	Power string
}

func (g Gun) Shot() {
	fmt.Printf("%v火力值%v", g.Name, g.Power)
}
