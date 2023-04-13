package mapdemo

type CustomMap struct {
	data map[string]Dog
}
type Dog struct {
	Name string
}

func GetDatas() CustomMap {

	return CustomMap{
		data: map[string]Dog{Name: ""}
	}
}
