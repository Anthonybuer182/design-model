package adapter

import "fmt"

type UsbToTypecAdapter struct {
}

func (t UsbToTypecAdapter) insert() {
	fmt.Println("adapter可直接接入")
	u := Usb{}
	u.usb()
}
