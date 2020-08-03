package pattern

/*
意图:
定义一系列的算法，把它们一个个封装起来，并且使它们可相互替换。

关键代码:
实现同一个接口

应用实例:

1、主题的更换，每个主题都是一种策略
2、旅行的出游方式，选择骑自行车、坐汽车，每一种旅行方式都是一个策略
*/

// 实现此接口，则为一个策略
type IStrategy interface {
	do(int, int) int
}

// 加法
type add struct {
}

// 实现加法
func (*add) do(a, b int) int {
	return a + b
}

// 减
type reduce struct {
}

// 实现减法
func (*reduce) do(a, b int) int {
	return a - b
}

// 具体策略的执行者
type Operator struct {
	strategy IStrategy
}

// 设置策略
func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

// 调用策略中的方法
func (operator Operator) calculate(a, b int) int {
	return operator.strategy.do(a, b)
}
