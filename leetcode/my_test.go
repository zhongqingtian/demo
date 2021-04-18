package leetcode

import (
	"fmt"
	"testing"
)

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

// 无重复数字全排列
func TestPermute(t *testing.T) {
	fmt.Println(Permute([]int{1, 2, 3}))
}

func TestPermuteUnique(t *testing.T) {
	fmt.Println(PermuteUnique([]int{1, 1, 2})) // 有重复数字全排列 [[1 1 2] [1 2 1] [2 1 1]]
}

func TestSubsets(t *testing.T) { // 子集
	fmt.Println(Subsets([]int{1, 2, 3}))
}

func TestSubsetsWithDup(t *testing.T) {
	fmt.Println(SubsetsWithDup([]int{1, 2, 2}))
}

// 组合 回溯 递归
func TestCombinationSum(t *testing.T) {
	fmt.Println(CombinationSum([]int{2, 3, 6, 7}, 7))
}

func TestCombinationSum2(t *testing.T) {
	fmt.Println(CombinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

// 打家劫舍---动态规划--房屋不是环
func TestRob(t *testing.T) {
	fmt.Println(Rob([]int{2, 7, 9, 3, 1}))
}

// 打家劫舍2---动态规划--房屋是环
func TestRob2(t *testing.T) {
	fmt.Println(Rob2([]int{2, 3, 2}))
}

// 打家劫舍3---动态规划--房屋是树状
func TestRob3(t *testing.T) {
	//fmt.Println(Rob3())
}

// 最大子序和
func TestMaxSubArray(t *testing.T) {
	fmt.Println(MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

// 乘积最大子数组
func TestMaxProduct(t *testing.T) {
	fmt.Println(MaxProduct([]int{2, 3, -2, 4}))
}

//三个数的最大乘积
func TestMaximumProduct(t *testing.T) {
	fmt.Println(MaximumProduct([]int{1, 2, 3, 4}))
}

func TestNextPermutation(t *testing.T) {
	a := []int{3, 2, 1}
	NextPermutation(a)
	fmt.Println(a)
}
func TestTrap(t *testing.T) {
	fmt.Println(Trap([]int{4, 2, 3}))
}
