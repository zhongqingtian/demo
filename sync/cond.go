package sync

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()

	//for !done { // 条件判断，不满足则进入等等
	c.Wait()
	fmt.Println("111")
	//}
	log.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	log.Println(name, "starts writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast()
	// c.Signal() // 从等待队列中移除第一个 goroutine 并把它唤醒
}

// https://studygolang.com/articles/34348?fr=sidebar
func CondRun() {
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)
	go read("reader4", cond) // 唤醒后，再次wait,可以再次唤醒，要不同协程，否则会阻塞
	write("writer", cond)

	time.Sleep(time.Second * 3)
}
