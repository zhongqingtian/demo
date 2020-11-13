package _defer

import "fmt"

func GetReturn() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println("defer 里面可以被改变", err)
		}
	}()

	err = fmt.Errorf("err change")
	return
}
