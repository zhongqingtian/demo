package leetcode

import "container/heap"

type TopList []int

func (t TopList) Len() int {
	return len(t)
}
func (t TopList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t TopList) Less(i, j int) bool {
	return t[i] < t[j]
}
func (t *TopList) Push(x interface{}) {
	*t = append(*t, x.(int))
}
func (t *TopList) Pop() interface{} {
	x := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return x
}

func FindKthLargest(nums []int, k int) int {
	m := make(TopList, 0)
	size := 0
	for i := range nums {
		if size < k {
			heap.Push(&m, nums[i])
			size++
		} else {
			if nums[i] > m[0] { //小顶堆 堆顶元素小于当前元素
				heap.Pop(&m)
				heap.Push(&m, nums[i])
			}
		}
	}
	return m[0]
}
