package adapter

import "fmt"

type Usb struct {
}

func (usb Usb) usb() {
	fmt.Println("usb对接adapter")
}
