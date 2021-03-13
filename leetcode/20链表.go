package leetcode

// 环形链表
// 给定一个链表，判断链表中是否有环。
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return false
		}
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 环形链表 II
/*
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
*/
func detectCycle(head *ListNode) *ListNode {
	s := head
	f := head
	for f != nil {
		if f.Next == nil {
			return nil
		}
		s, f = s.Next, f.Next.Next
		if s == f {
			f = head
			for {
				if s == f {
					return s
				}
				s, f = s.Next, f.Next
			}
		}
	}
	return nil
}

// 链表相交
/*给定两个（单向）链表，判定它们是否相交并返回交点。请注意相交的定义基于节点的引用，而不是基于节点的值。换句话说，如果一个链表的第k个节点与
另一个链表的第j个节点是同一节点（引用完全相同），则这两个链表相交。
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	temp := make(map[*ListNode]int)

	for headA != nil {
		temp[headA] = 1
		headA = headA.Next
	} // 循环进入map

	for headB != nil {
		if temp[headB] == 1 {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// 两个链表的第一个公共节点
/*
解题思路：
我们使用两个指针 node1，node2 分别指向两个链表 headA，headB 的头结点，然后同时分别逐结点遍历，当 node1 到达链表 headA 的末尾时，重新定位到链表 headB 的头结点；当 node2 到达链表 headB 的末尾时，重新定位到链表 headA 的头结点。

这样，当它们相遇时，所指向的结点就是第一个公共结点。
*/
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	node1, node2 := headA, headB

	for node1 != node2 {
		if node1 == nil {
			node1 = headB
		} else {
			node1 = node1.Next
		}
		if node2 == nil {
			node2 = headA
		} else {
			node2 = node2.Next
		}
	}
	return node1
}
