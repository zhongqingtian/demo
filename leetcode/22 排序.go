package leetcode

import "sort"

// 合并区间
/*
数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。



示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := make([][]int, 0)
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		m := merged[len(merged)-1]
		c := intervals[i]
		if m[1] < c[0] {
			merged = append(merged, c)
		}

		if m[1] > c[1] {
			m[1] = c[1]
		}
	}

	return merged
}


// 排序链表
//
//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // 递归的出口，不用排序 直接返回
		return head
	}
	slow, fast := head, head // 快慢指针
	var preSlow *ListNode    // 保存slow的前一个结点
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next      // 慢指针走一步
		fast = fast.Next.Next // 快指针走两步
	}
	preSlow.Next = nil  // 断开，分成两链
	l := sortList(head) // 已排序的左链
	r := sortList(slow) // 已排序的右链
	return mergeList(l, r) // 合并已排序的左右链，一层层向上返回
}

func mergeList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}   // 虚拟头结点
	prev := dummy                // 用prev去扫，先指向dummy
	for l1 != nil && l2 != nil { // l1 l2 都存在
		if l1.Val < l2.Val {   // l1值较小
			prev.Next = l1 // prev.Next指向l1
			l1 = l1.Next   // 考察l1的下一个结点
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next // prev.Next确定了，prev指针推进
	}
	if l1 != nil {    // l1存在，l2不存在，让prev.Next指向l1
		prev.Next = l1
	}
	if l2 != nil {
		prev.Next = l2
	}
	return dummy.Next // 真实头结点
}
/*
时间复杂度是 O(nlogn)
空间复杂度是 O(logn)

我没写出来空间复杂度是 O(1) 的迭代版本，掌握递归版的其实也够了。

mergeList 的代码也可以写成递归版的：*/
/*
func mergeList(l1, l2 *ListNode) *ListNode {
	if l1==nil {
		return l2
	}
	if l2==nil {
		return l1
	}
	if l1.Val > l2.Val{
		l2.Next = mergeList(l1, l2.Next)
		return l2
	}
	l1.Next = mergeList(l2, l1.Next)
	return l1
}
*/