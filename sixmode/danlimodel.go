package sixmode

import (
	"fmt"
	"sync"
)

type signleton struct {
}

var (
	instance *signleton
	once     sync.Once
)

func Getinstance() *signleton {
	once.Do(func() {
		instance = &signleton{}
		fmt.Println("&&&&:", &instance)
	})
	fmt.Println("&&&&222:", &instance)
	return instance
}

/*
var mu sync.Mutex

type singleton struct {
}

var instance *singleton

func Getinstance() *singleton {
	if instance == nil {
		mu.Lock()
		if instance == nil {

			instance = &singleton{}
		}
		mu.Unlock()
	}

	return instance
}

*/
