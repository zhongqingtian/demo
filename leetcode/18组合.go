package leetcode

import "sort"

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
*/
func CombinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)

	var dfs func(start int, path []int, sum int)

	dfs = func(start int, path []int, sum int) {
		if sum >= target {
			if sum == target {
				temp := make([]int, len(path))
				copy(temp, path)
				ans = append(ans, temp)
			}
			return
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			dfs(i, path, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{}, 0)
	return ans
}

/*
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。
*/

func CombinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(candidates)
	var dfs func(lastNum []int, path []int, sum int)
	dfs = func(lastNum []int, path []int, sum int) {
		if sum >= target {
			if sum == target {
				temp := make([]int, len(path))
				copy(temp, path)
				ans = append(ans, temp)
			}
			return
		}

		for i, n := range lastNum {
			if i > 0 && lastNum[i-1] == lastNum[i] {
				continue
			}
			path = append(path, n)
			dfs(lastNum[i+1:], path, sum+lastNum[i])
			path = path[:len(path)-1]
		}
	}

	dfs(candidates, []int{}, 0)
	return ans
}
