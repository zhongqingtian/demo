package helpers

import "testing"

//单元测试
func TestValidateMobile(t *testing.T) {
	result := ValidateMobile("190")
	if result != false {
		t.Error("the result is wrong")
	}
}

// 性能测试
func BenchmarkValidateMobile(b *testing.B) {
	for i := 0;i < b.N ;i++  {
		ValidateMobile("130")
	}
}