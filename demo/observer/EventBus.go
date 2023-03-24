package observer

import "fmt"

type EventBus struct {
	observers []Observer
}

func NewEventBus() *EventBus {
	e := &EventBus{}
	return e
}
func (e *EventBus) Register(o Observer) {
	e.observers = append(e.observers, o)
	fmt.Println(len(e.observers))
}
func (e *EventBus) UnRegister(o Observer) []Observer {
	for i, v := range e.observers {
		if v.getId() == o.getId() {
			// 将要剔除的元素和最后一个元素互调 然后长度剪一
			e.observers[len(e.observers)-1], e.observers[i] = e.observers[i], e.observers[len(e.observers)-1]
			return e.observers[:len(e.observers)-1]
		}
	}
	return e.observers
}

func (e *EventBus) PublishAll(msg string) []Observer {
	for _, v := range e.observers {
		v.subscribe(msg)
	}
	return e.observers
}
