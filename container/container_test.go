package container

import (
	"container/heap"
	"container/list"
	"container/ring"
	"testing"
)

func TestIntHeap_Pop(t *testing.T) {
	h := &IntHeap{2, 1, 6, 4}
	heap.Init(h)
	heap.Push(h, 3)
	t.Log(h)
	heap.Pop(h)
}

func TestUseList(t *testing.T) {
	UseList()
	list1 := list.New()
	list1.PushBack(1)        // 往链表后面添加元素
	e1 := list1.PushFront(2) // 往链表前面添加元素  这时 2，1

	list1.InsertAfter(3, e1) // 在2后面添加3 这时 2，3，1
	t.Log("len: ", list1.Len())
	t.Log("value: ", list1.Front().Value)
	t.Log("value: ", list1.Front().Next().Value)
	t.Log("value: ", list1.Front().Next().Next().Value)
	/* 结果
	   container_test.go:24: len:  3
	   container_test.go:25: value:  2
	   container_test.go:26: value:  3
	   container_test.go:27: value:  1
	*/

	// 环
	ring := ring.New(3) // 创建一个包含n个元素的环

	for i := 1; i <= 3; i++ {
		ring.Value = i
		ring = ring.Next()
	}

	// 计算 1+2+3
	s := 0
	ring.Do(func(p interface{}) { // Do()方法迭代一个环所有元素，然后通过传输方法类型叠加
		s += p.(int)
	})
	t.Log("sum is", s)
}

func TestSlic_Do(t *testing.T) {
	s := &Slic{2, 6, 2}
	sum := 0
	s.Do(func(p interface{}) { // 参数是一个完整的参数
		sum += p.(int)
	})

	t.Log(sum)
}
