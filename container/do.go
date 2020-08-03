package container

type Slic []int

func (s Slic) Do(f func(v interface{})) { // 只把值传给接口方法，具体方法实现
	l := len(s)
	for i := 0; i < l; i++ {
		f(s[i])
	}
}
