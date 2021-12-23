package leetcode

import (
	"bytes"
	"sort"
)

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

/*
整数转罗马数字
*/
func intToRoman(num int) string {
	/**
	I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
	X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
	C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
	*/
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	ss := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var i = 0
	var buffer bytes.Buffer
	for num > 0 && i < len(nums) {
		if num > nums[i] {
			buffer.WriteString(ss[i])
			num = num - nums[i]
		} else {
			i++
		}
	}
	return buffer.String()
}

/*
15 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	var addNum int
    sort.Ints(nums)
	var start, index, right int
	for index = 1; index < len(nums)-1; index++ {
		start, right = 0, len(nums)-1
		// index 去重
		if index > 1 && nums[index-1] == nums[index] {
			start = index - 1
		}
		for start < index && index < right {
			addNum = nums[start] + nums[index] + nums[right]
			// start 去重
			if start > 1 && nums[start-1] == nums[start] {
				start++
				continue
			}
			// right 去重
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			if addNum == 0 {
				res = append(res, []int{nums[start], nums[index], nums[right]})
			   start++
			   right--
			} else if addNum > 0 {
				right--
			} else {
				start++
			}
		}
	}
	return res
}


/*func threeSum(nums []int) [][]int {
	var saveArray [][]int
	var addNum int
	sort.Ints(nums)

	var start,index,right int

	for index=1;index<len(nums)-1;index++ {
		start,right = 0,len(nums)-1
		// index 去重
		if index>1&&nums[index]==nums[index-1]{
			start = index-1
		}

		for start<index && index <right {
			addNum = nums[start] +  nums[index] + nums[right]

			// 左边去重
			if start>0 && nums[start]==nums[start-1] {
				start++
				continue
			}

			// 右边去重
			if right<len(nums)-1&& nums[right]==nums[right+1] {
				right--
				continue
			}

			if addNum==0 {
				saveArray = append(saveArray,[]int{nums[start],nums[index],nums[right]})
				start++
				right--
			} else if addNum > 0{
				right--
			} else {
				start++
			}


		}
	}
	return saveArray
}
*/