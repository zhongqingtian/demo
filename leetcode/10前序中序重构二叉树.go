package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
*/
// 一个包可以有多个init()方法
func init() {

}

func init() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 {
		return nil
	}

	// 根节点
	rootVal := preorder[0]

	root := &TreeNode{
		Val: rootVal,
	}
	var i = 0
	for ; i < len(inorder); i++ {
		if inorder[i] == rootVal { // 找到根节点
			break
		}
	}

	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return root
}
