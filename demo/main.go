package main

import (
	"design-model/demo/adapter"
	"design-model/demo/build"
	"design-model/demo/chain"
	"design-model/demo/customerr"
	"design-model/demo/factory"
	"design-model/demo/filter"
	"design-model/demo/mapdemo"
	"design-model/demo/observer"
	"design-model/demo/recursion"
	"design-model/demo/signle"
	"design-model/demo/strtegy"
	"fmt"
)

func testSingle() {
	for i := 0; i < 10; i++ {
		go signle.GetInstance()
	}
	fmt.Scanln()
}
func testFactory() {
	gun := factory.CreateGun("ak")
	gun.Shot()
}
func testAdapter() {
	client := adapter.MacClient{}
	typec := adapter.MacTypec{}
	client.InsertCharge(typec)
	adapter := adapter.UsbToTypecAdapter{}
	client.InsertCharge(adapter)
}
func testBuild() {
	house := build.NewBuilder().SetDoor("门").SetFloor("地板").SetWindow("窗户").Build()
	fmt.Printf("house: %v  type: %T", house, house)
}
func testChain() {
	c := &chain.Cashier{}

	m := &chain.Medical{}
	m.Next(c)

	d := &chain.Doctor{}
	d.Next(m)

	r := &chain.Reception{}
	r.Next(d)

	p := &chain.Patient{}
	r.Execute(p)

}
func testObserver() {
	bus := observer.NewEventBus()
	c := observer.Customer{}
	bus.Register(c)
	bus.PublishAll("发布数据")

}

type (
	Person = filter.Person
)

func testFilter() {
	p := []Person{}
	p = append(p, Person{Name: "张三", Age: 15, Gender: 1})
	p = append(p, Person{Name: "李四", Age: 20, Gender: 0})
	p = append(p, Person{Name: "王五", Age: 30, Gender: 1})
	p = append(p, Person{Name: "赵六", Age: 35, Gender: 0})
	newP := filter.NewFilterClient().SetPersons(p).AddFilter(filter.AgeFilter{}).AddFilter(filter.GenderFilter{}).GetFilterPersons()
	fmt.Println(newP)
}
func testCount(num1 float64, num2 float64) {
	add := strtegy.Count(strtegy.Add{}, num1, num2)
	sub := strtegy.Count(strtegy.Sub{}, num1, num2)
	fmt.Printf("%v+%v=%v\n", num1, num2, add)
	fmt.Printf("%v-%v=%v", num1, num2, sub)
}
func testRecursion(num int) {
	fmt.Println(recursion.Factorial(num))
}
func testCustomerr() {
	err := customerr.OpenFile("")
	fmt.Println(err)
}
func testMap() {
	datas := mapdemo.GetDatas()
	fmt.Println(datas)
}
func main() {
	testMap()
}
