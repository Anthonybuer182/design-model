package build

type IBuilder interface {
	setDoor()
	setWindow()
	setFloor()
	build()
}
