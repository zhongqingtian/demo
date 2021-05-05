package leetcode

/*
盛最多水的容器
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
*/
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	maxArea := 0
	for i != j {
		hi, hj := height[i], height[j]
		s := (j - i) * min(hi, hj)
		if s > maxArea {
			maxArea = s
		}
		if hi < hj {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

/*
接雨水接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

func trap(height []int) int {
	if len(height) < 0 {
		return 0
	}
	left, right := 0, len(height)-1
	var res int
	leftMax := 0
	rightMax := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right]>rightMax {
				rightMax=height[right]
			}else {
				res += rightMax-height[right]
			}
			right--
		}
	}
	return res
}

