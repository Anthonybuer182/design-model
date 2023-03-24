package build

type Builder struct {
	door   string
	window string
	floor  string
}

func NewBuilder() *Builder {
	return &Builder{}
}
func (b *Builder) SetDoor(door string) *Builder {
	b.door = door
	return b
}
func (b *Builder) SetWindow(window string) *Builder {
	b.window = window
	return b
}
func (b *Builder) SetFloor(floor string) *Builder {
	b.floor = floor
	return b
}
func (b *Builder) Build() House {
	return House{door: b.door, window: b.window, floor: b.floor}
}
