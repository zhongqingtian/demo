package container

import "sort"

/*这里的堆使用的数据结构是最小二叉树，即根节点比左边子树和右边子树的所有值都小。
go 的堆包只是实现了一个接口，我们看下它的定义：
*/
type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}

/*可以看出，这个堆结构继承自 sort.Interface, 回顾下 sort.Interface，它需要实现三个方法

Len() int
Less(i, j int) bool
Swap(i, j int)
加上堆接口定义的两个方法

Push(x interface{})
Pop() interface{}
就是说你定义了一个堆，就要实现五个方法
*/

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
