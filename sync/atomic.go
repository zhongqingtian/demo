package sync

import (
	"github.com/sirupsen/logrus"
	"sync/atomic"
	"time"
)

func TAddInt() {
	var i32 int32
	logrus.Info("======old i32 value=====")
	logrus.Info(i32)
	// 第一个参数值必须是一个指针类型的值,因为该函数需要获得被操作值在内存中的存放位置,以便施加特殊的CPU指令
	// 结束时会返回原子操作后的新值
	for i := 0; i < 1000; i++ {
		go func() {
			//atomic.AddInt32(&i32, 1)
			i32 = i32 + 1
			logrus.Info("read = ", i32)
		}()
	}
	time.Sleep(3 * time.Second)
	newI32 := atomic.AddInt32(&i32, 6)
	logrus.Info("======new i32 value ======")
	logrus.Info(i32)
	logrus.Info(newI32)

}

var value int32

// 比较然后替换: 不断尝试原子地址更新value,直到操作成功为止
func CAS(delta int32) {
	// 在被操作值频繁变更的情况下，CAS操作并不会那么容易成功
	// so 不得不利用for 循坏进行多次尝试
	for i := 0; i < 100; i++ {
		go func() {
			value = int32(i)
		}()
	}
	for {
		v := value
		logrus.Info("before = ", value)
		if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
			logrus.Info("after = ", value)
			// 在函数结果值为true时，退出循坏
			break
		}
		//操作失败的缘由总会是value的旧值已不与v的值相等了.
		//CAS操作虽然不会让某个Goroutine阻塞在某条语句上,但是仍可能会使流产的执行暂时停一下,不过时间大都极其短暂.
	}
}

func LoadInt(delta int32) {
	for {
		//v := value
		//在进行读取value的操作的过程中,其他对此值的读写操作是可以被同时进行的,那么这个读操作很可能会读取到一个只被修改了一半的数据.
		//因此我们要使用载入
		v := atomic.LoadInt32(&value)
		if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
			//在函数的结果值为true时,退出循环
			break
		}
		//操作失败的缘由总会是value的旧值已不与v的值相等了.
		//CAS操作虽然不会让某个Goroutine阻塞在某条语句上,但是仍可能会使流产的执行暂时停一下,不过时间大都极其短暂.
	}
}
