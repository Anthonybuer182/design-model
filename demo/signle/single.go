package signle

import (
	"fmt"
	"sync"
)

type signle struct{}

var lock = &sync.Mutex{}
var singleInstance *signle

func GetInstance() *signle {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("created0")
			singleInstance = &signle{}
		} else {
			fmt.Println("created")
		}
	} else {
		fmt.Println("created1")
	}
	return singleInstance
}
