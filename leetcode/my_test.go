package leetcode

import "testing"

func TestLetterCombinations(t *testing.T) {
	//t.Log(LetterCombinations("23"))
	// t.Log(FindKthLargest([]int{3, 2, 1, 5, 6, 4}, 5))
	var a interface{}
	a = 2
	k := a.(int)
	t.Log(k)
}

func TestGenerateParenthesis(t *testing.T) {
	t.Log(GenerateParenthesis(3))
}

func TestIncreasingBST(t *testing.T) {
	node := &TreeNode{Val: 10, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}}}
	res := IncreasingBST(node)
	for res != nil {
		t.Log(res.Val)
		res = res.Right
	}
}
