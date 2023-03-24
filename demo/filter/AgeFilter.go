package filter

type AgeFilter struct {
}

// 拦截年龄大于20岁的
func (af AgeFilter) Intercept(p Person) bool {
	return p.Age > 20
}
