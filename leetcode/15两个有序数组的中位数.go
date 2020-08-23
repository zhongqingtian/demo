package leetcode

/*
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。

请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

如何把时间复杂度降低到 O(\log(m+n))O(log(m+n)) 呢？如果对时间复杂度的要求有 \loglog，通常都需要用到二分查找，这道题也可以通过二分查找实现。
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
	return 0
}

// 找两个数组，第k个最小的数
func getKthElement(nums1, nums2 []int, k int) int {
	indexSum1 := 0 // 记录两个数组的下标位置
	indexSum2 := 0
	for {
		if indexSum1 == len(nums1) { // 已经取出sum1中的元素，说明数字都小
			return nums2[indexSum2+k-1]
		}
		if indexSum2 == len(nums2) {
			return nums1[indexSum1+k-1]
		}
		if k == 1 {
			return min(nums1[indexSum1], nums2[indexSum2]) // 最后一次，两个最小数比较
		}
		half := k / 2

		newIndex1 := min(indexSum1+half, len(nums1)) - 1
		newIndex2 := min(indexSum2+half, len(nums1)) - 1
		provt1 := nums1[newIndex1]
		provt2 := nums1[newIndex2]

		if provt1 < provt2 {
			k -= newIndex1 - indexSum1 - 1
			indexSum1 += newIndex1 + 1 // sum1 下移下标一次
		} else {
			k -= newIndex2 - indexSum2 - 1
			indexSum2 += newIndex2 + 1
		}
	}
}
