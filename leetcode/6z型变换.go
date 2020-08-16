package leetcode

import "math"

/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
*/

func Zconvert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	groupLen := numRows*2 - 2
	// 分成多少个组，组的数目
	groupNum := int(math.Ceil(float64(len(s)) / float64(groupLen)))
	var ansSring []byte
	for i := 0; i < numRows; i++ { //一行一行取数据
		for j := 0; j < groupNum; j++ { // 一组一组取数据
			charIndex := j*groupLen + i
			if charIndex >= len(s) {
				continue
			}
			ansSring = append(ansSring, s[charIndex])
			if i != 0 && i != numRows-1 {
				charIndex = (j+1)*groupLen - i
				if charIndex < len(s) {
					ansSring = append(ansSring, s[charIndex])
				}
			}

		}
	}

	return string(ansSring)
}

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	groupLen := numRows*2 - 2
	groupNum := int(math.Ceil(float64(len(s)) / float64(groupLen)))
	var ansString []byte

	for i := 0; i < numRows; i++ {
		//计算第 i 行字符串
		for j := 0; j < groupNum; j++ {
			//计算第 j 组字符串
			charIndex := groupLen*j + i
			if charIndex >= len(s) {
				continue
			}
			ansString = append(ansString, s[charIndex])
			if i != 0 && i != numRows-1 {
				charIndex = groupLen*(j+1) - i
				if charIndex < len(s) {
					ansString = append(ansString, s[charIndex])
				}
			}
		}

	}
	return string(ansString)
}
