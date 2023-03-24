package filter

// go 接口中范型不好使
type IFilter interface {
	Intercept(p Person) bool
}
