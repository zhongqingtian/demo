package leetcode

/*
打家劫舍 I

你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
*/

func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ { // 状态转移方程
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(dp)-1]
}

/*
打家劫舍 II
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，
这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在
同一晚上被小偷闯入，系统会自动报警 。
输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
*/

func Rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// 包含第一家,不包含最后一家
	robSum1 := Rob(nums[:len(nums)-1])
	// 不包含第一家，包含最后一家
	robSum2 := Rob(nums[1:])
	return max(robSum1, robSum2)
}

//  打家劫舍 III
/*
在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。

计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

示例 1:

输入: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

输出: 7
解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
*/

/*  type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}*/

func Rob3(root *TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)

	selected := node.Val + l[1] + r[1]               // 当前节点选择，则两个字节点都不能被选择
	NotSelected := max(l[0], l[1]) + max(r[1], r[0]) // 当前节点不选择，子节点可以被选或者不选
	return []int{selected, NotSelected}
}
