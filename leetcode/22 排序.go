package leetcode

import "sort"

// 合并区间
/*
数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。



示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := make([][]int, 0)
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		m := merged[len(merged)-1]
		c := intervals[i]
		if m[1] < c[0] {
			merged = append(merged, c)
		}

		if m[1] > c[1] {
			m[1] = c[1]
		}
	}

	return merged
}
