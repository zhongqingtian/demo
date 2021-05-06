package leetcode

import "sort"

/*
最大子序和
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。



示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/

func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	var max int = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] { // 加上的和是否比当前数字大
			nums[i] = nums[i-1] + nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 乘积最大子数组
/*
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
*/

func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	maxF := make([]int, len(nums)) // 存过程中的最大值
	minF := make([]int, len(nums)) // 存过程中的最小值
	maxF[0], minF[0] = nums[0], nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		maxF[i] = max3(maxF[i-1]*nums[i], minF[i-1]*nums[i], nums[i])
		minF[i] = min3(maxF[i-1]*nums[i], minF[i-1]*nums[i], nums[i])
		if maxF[i] > max {
			max = maxF[i]
		}
	}
	return max
}

func max3(a, b, c int) int {
	var t int = 0
	if a >= b {
		t = a
	} else {
		t = b
	}
	if t > c {
		return t
	}
	return c
}

func min3(a, b, c int) int {
	var t int = 0
	if a >= b {
		t = b
	} else {
		t = a
	}
	if t > c {
		return c
	}
	return t
}

/*
三个数的最大乘积
给你一个整型数组 nums ，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

输入：nums = [1,2,3,4]
输出：24
*/
func MaximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return max(nums[0]*nums[1]*nums[n-1], nums[n-3]*nums[n-2]*nums[n-1])
}
