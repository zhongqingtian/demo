package demo

import (
	"fmt"
	"testing"
)

// 没有泛型之前
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 这个时候我需要增加一个获取float64的最大值，那么我们就要新增一个函数叫maxFloat64：
func maxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func TestMax(t *testing.T) {
	fmt.Println(maxInt(32, 64))
}


