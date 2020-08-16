package leetcode

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
 思路：动态规划
解题思路
状态定义
dp[l,r]：字符串s从索引l到r的子串是否是回文串
true： s[l,r] 是回文串
false：s[l,r] 不是回文串
转移方程
dp[l][r] = dp[l+1][r-1] && s[l] == s[r]
s[l] == s[r]：说明当前中心可以继续扩张，进而有可能扩大回文串的长度
dp[l+1][r-1]：true
说明s[l,r]的**子串s[l+1][r-1]**也是回文串
说明，l是从最大值开始遍历的，r是从最小值开始遍历的
特殊情况
l - r <= 2：2表示aba类型；1表示aa类型；0表示a类型
总结
dp[l][r] = s[l] == s[r] && ( dp[l+1][r-1] || r - l <= 2)

*/

func LongestPalindrome(s string) string {
	lenth := len(s)
	if lenth <= 1 {
		return s
	}

	dp := make([][]bool, lenth)
	start := 0
	maxlen := 1

	for r := 0; r < lenth; r++ {
		dp[r] = make([]bool, lenth)
		dp[r][r] = true
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
			} else {
				dp[l][r] = false
			}

			if dp[l][r] {
				curlen := r - l + 1
				if curlen > maxlen {
					maxlen = curlen
					start = l
				}
			}
		}
	}
	return s[start : start+maxlen]
}
