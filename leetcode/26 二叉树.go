package leetcode


func verifyPostorder(postorder []int) bool {

	return false
}

/*后序遍历*/
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack []*TreeNode
	stack = append(stack, root)
	var ret []int
	for len(stack) > 0 {
		getStack := stack[len(stack)-1]
		// 判断栈首是否为叶子节点,若是,则出栈,取出叶子节点值,加入到ret等等返回
		if getStack.Left == nil && getStack.Right == nil {
			ret = append(ret, getStack.Val)
			stack = stack[:len(stack)-1]
			continue
		}

		// 由于前一个判断已经确定非叶子节点,则将其左右子节点加入到栈中,并将其左右子节点指针赋空,避免进入循环加入
		if getStack.Right != nil {
			stack = append(stack, getStack.Right)
			getStack.Right = nil
		}
		if getStack.Left != nil {
			stack = append(stack, getStack.Left)
			getStack.Left = nil
		}
	}
	return ret
}

/*中序遍历*/
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for len(stack) > 0 || root !=nil {
		for root != nil {
			stack = append(stack,root)// 入栈
			root = root.Left
		}

		// 出栈
		index := len(stack)-1
		res = append(res,stack[index].Val)

		root = stack[index].Right //右节点会进入下次循环，如果 =nil，继续出栈
		stack = stack[:index]
	}

	return res
}

/*后序遍历*/
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for 0 < len(stack) || root != nil { //root != nil 只为了第一次root判断，必须放最后
		for root != nil {
			res = append(res, root.Val)       //前序输出
			stack = append(stack, root.Right) //右节点 入栈
			root = root.Left                  //移至最左
		}
		index := len(stack) - 1 //栈顶
		root = stack[index]     //出栈
		stack = stack[:index]
	}
	return res

}