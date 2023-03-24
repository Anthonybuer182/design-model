package filter

type GenderFilter struct {
}

// 拦截女性 男性为1 女性为0
func (af GenderFilter) Intercept(p Person) bool {
	return p.Gender == 0
}
