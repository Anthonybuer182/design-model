package filter

type FilterClient struct {
	filters []IFilter
	persons []Person
}

func NewFilterClient() *FilterClient {
	fc := &FilterClient{}
	return fc
}
func (fc *FilterClient) AddFilter(f IFilter) *FilterClient {
	fc.filters = append(fc.filters, f)
	return fc
}
func (fc *FilterClient) SetPersons(persons []Person) *FilterClient {
	fc.persons = persons
	return fc
}
func (fc *FilterClient) GetFilterPersons() []Person {
	newPersons := []Person{}
	for _, p := range fc.persons {
		var intercept bool = false
		for _, f := range fc.filters {
			intercept = f.Intercept(p) || intercept
		}
		if !intercept {
			newPersons = append(newPersons, p)
		}
	}
	return newPersons
}
