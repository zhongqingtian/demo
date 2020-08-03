package project

import (
	"fmt"
	"sync"
)

type WorkPool struct {
	wg sync.WaitGroup
}

func (wp *WorkPool) Run(function func()) {
	wp.wg.Add(1)
	go func() {
		function() // 执行并发的方法
		wp.wg.Done()
	}()
}

func (wp *WorkPool) Wait() {
	wp.wg.Wait()
}

func MakeWaitGroup() {
	var syn = WorkPool{}
	syn.Run(
		func() {
			for i := 1000; i > 0; i-- {
				fmt.Println(i)
			}
		})
	syn.Wait() // 等待并发协程执行结束再最后输出结束语句
	fmt.Println("game over!")
}
