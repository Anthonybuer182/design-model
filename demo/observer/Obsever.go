package observer

type Observer interface {
	subscribe(string)
	getId() int
}
