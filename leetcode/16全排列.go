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
