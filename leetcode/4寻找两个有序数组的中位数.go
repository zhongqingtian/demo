package leetcode

/*给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
*/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n { // 更换数组，确保 nums1 长度短的
		t := nums1
		nums1 = nums2
		nums2 = t
		m = len(nums1)
		n = len(nums2)
	}
	iMin, iMax, halfLen := 0, m, (m+n+1)/2
	for iMin <= iMax {
		i := (iMax + iMin) / 2
		j := halfLen - i
		if i < iMax && nums2[j-1] > nums1[i] {
			iMax = i - 1
		} else if i > iMin && nums2[j] < nums1[i-1] {
			iMax = i - 1
		} else {

		}
	}

	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
