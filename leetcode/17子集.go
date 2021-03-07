package leetcode

import "sort"

//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
func Subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int(nil))

	var dfs func([]int, []int)
	dfs = func(lastNums []int, path []int) {
		if len(lastNums) == 0 {
			return
		}
		for i, n := range lastNums {
			path = append(path, n)
			tempPath := make([]int, len(path))
			copy(tempPath, path)
			res = append(res, tempPath)
			dfs(lastNums[i+1:], path)
			path = path[:len(path)-1]
		}
	}
	dfs(nums, []int{})
	return res
}

/*
输入: [1,2,2]
输出:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
*/
// 有重复数字的子集
func SubsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int(nil))
	sort.Ints(nums)

	var dfs func(lastNums []int, path []int)
	dfs = func(lastNums []int, path []int) {
		if len(lastNums) == 0 {
			return
		}
		for i, num := range lastNums {
			if i > 0 && lastNums[i-1] == lastNums[i] {
				continue
			}
			path = append(path, num)
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			dfs(lastNums[i+1:], path)
			path = path[:len(path)-1]
		}
	}
	dfs(nums, []int{})
	return res
}
