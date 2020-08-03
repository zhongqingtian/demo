package pattern

import "sync"

var (
	goInstance *Instance
	once       sync.Once
)

type Instance struct {
	Name string
}

// go 特有的单例模式
func GoInstance(name string) *Instance {
	if goInstance == nil {
		once.Do(func() {
			goInstance = &Instance{Name: name}
		})
	}
	return goInstance
}
