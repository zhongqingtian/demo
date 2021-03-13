package leetcode

import "sort"

// []{1,1,2} 有重复数字全排列
func PermuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || (i > 0 && !vis[i-1] && nums[i-1] == nums[i]) { // !vis[i-1] 前一个刚刚回退，nums[i-1] == nums[i]和前一位相同
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}

// 无重复数字全排列
func Permute(nums []int) [][]int {
	res := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for _, n := range nums {
			if visited[n] {
				continue
			}
			path = append(path, n)
			visited[n] = true
			dfs(path)
			path = path[:len(path)-1]
			visited[n] = false
		}
	}

	dfs([]int{})
	return res
}

// 下一个排列
/*
1234753
1235347
*/
func NextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i] // 交换位置
	}
	reverse(nums[i+1:]) // 把后面 列数字反转排序
}
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

/*
接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
输入：height = [4,2,0,3,2,5]
输出：9
*/
func Trap(height []int) int {  // 双指针
	if len(height) < 0 {
		return 0
	}
	var res int
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left] // 设置左边高的柱子
			} else {
				res += leftMax - height[left] // 左边必有池子挡水
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right--
		}
	}
	return res
}
