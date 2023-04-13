package decorator

type IPizza interface {
	GetPrice() int
}
type Chicago struct {
	Price int
}

var _ IPizza = (*Chicago)(nil)

func (c *Chicago) GetPrice() int {
	return c.Price
}

type Tomato struct {
	Price int
	Pizza IPizza
}

func (c *Tomato) GetPrice() int {
	return c.Pizza.GetPrice() + c.Price
}
