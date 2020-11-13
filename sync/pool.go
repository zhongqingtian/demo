package sync

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)

// 用来保存和复用临时对象，以减少内存分配，降低CG压力
// 如果 Pool 为空，则调用 New 返回一个新创建的对象
/*
1、如果没有设置 New，则返回 nil。
2、还有一个重要的特性是，放进 Pool 中的对象，会在说不准什么时候被回收掉。
3、所以如果事先 Put 进去 100 个对象，下次 Get 的时候发现 Pool 是空也是有可能的。
不过这个特性的一个好处就在于不用担心Pool 会一直增长，因为 Go 已经帮你在 Pool 中做了回收机制。
4、这个清理过程是在每次垃圾回收之前做的。垃圾回收是固定两分钟触发一次
5、 pool 创建的时候是不能指定大小的，所有 sync.Pool 的缓存对象数量是没有限制的（只受限于内存）

*/

func PoolGet() {
	// 建立对象
	initNewFuc := func() interface{} { return "Hello, BeiJing" } // 初始化new 对象的方法
	var pipe = &sync.Pool{New: initNewFuc}

	// 准备放入的字符串
	val := "Hello,World!"

	// 放入
	pipe.Put(val)
	// 取出
	log.Println(pipe.Get())
	// 再取就没有了,会自动调用NEW
	log.Println(pipe.Get())

}

/*
 pool 包在 init 的时候注册了一个 poolCleanup 函数，它会清除所有的 pool 里面的所有缓存的对象
该函数注册进去之后会在每次 gc 之前都会调用，因此 sync.Pool 缓存的期限只是两次 gc 之间这段时间
*/
func PoolGc() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	a := p.Get().(int)
	p.Put(1)
	runtime.GC()
	b := p.Get().(int)
	fmt.Println(a, b) // 输出是0，0 而不是 0，1
}
