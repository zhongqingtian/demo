package leetcode

import "strconv"

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数
*/

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	len := len(s)
	for i, _ := range s {
		if s[i] != s[len-i-1] {
			return false
		}
	}

	return true
}

/*
最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。ababac  => ababa
*/
func longestPalindrome(s string) string {
	length := len(s)
	dp := make([][]bool, length)
	start := 0
	maxLen := 1
	for r := 0; r < length; r++ {
		dp[r] = make([]bool, length)
		dp[r][r] = true
		for l := 0; l < r; l++ {
			if (s[l] == s[r]) && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
			} else {
				dp[l][r] =false
			}
			if dp[l][r] {
                curLen := r -l + 1
                if curLen>maxLen{
                	start = l
                	maxLen = curLen
				}
			}
		}
	}
	return s[start:start+maxLen]
}
