package leetcode

/*
https://leetcode-cn.com/problems/majority-element-ii/
229 求众数 II
给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。
*/
func majorityElement(nums []int) []int {
	res := make([]int, 0)
	if nums == nil || len(nums) == 0 {
		return nums
	}
	can1 := nums[0]
	can2 := nums[0]
	count1 := 0
	count2 := 0
	//摩尔投票法
	// 配对阶段
	for _, num := range nums {
		if can1 == num {
			count1++
			continue
		}
		if can2 == num {
			count2++
			continue
		}

		if count1 == 0 {
			can1 = num
			count1++
			continue
		}
		if count2 == 0 {
			can2 = num
			count2++
			continue
		}

		count1--
		count2--
	}
	// 计数
	count1 = 0
	count2 = 0
	for _, num := range nums {
		if can1 == num {
			count1++
			continue
		}
		if can2 == num {
			count2++
			continue
		}
	}
	if count1 > len(nums)/3 {
		res = append(res, can1)
	}
	if count2 > len(nums)/3 {
		res = append(res, can2)
	}
	return res
}
