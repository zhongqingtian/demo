package leetcode

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
示例：
输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
*/

/*
解题思路
回溯算法，回溯跳出条件就是左右括号都已经排完的情况。
括号成对存在，先有左括号再有右括号，所以只有右括号的数量小于左括号才进行右括号的添加。
最后如果右括号的数量等于0，表示右括号已经排完了，同时意味着左括号也排完了。
*/
func GenerateParenthesis(n int) []string {
	res := new([]string)
	backTracking(n, n, "", res)
	return *res // *标识去指针，取值
}

func backTracking(left, right int, str string, res *[]string) {
	if right == 0 { // 右边括号也拼接完了
		*res = append(*res, str)
	}

	if left > 0 {
		backTracking(left-1, right, str+"(", res)
	}

	if right > left {
		backTracking(left, right-1, str+")", res)
	}
}
