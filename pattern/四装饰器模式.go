package pattern

import (
	"github.com/sirupsen/logrus"
)

/*意图:
装饰器模式动态地将责任附加到对象上。若要扩展功能，装饰者提供了比继承更有弹性的替代方案。

关键代码:
装饰器和被装饰对象实现同一个接口，装饰器中使用了被装饰对象

应用实例:
JAVA 中的 IO 流。
*/
type Person interface {
	cost() int
	show()
}

// 被装饰对象
type laowang struct {
}

func (*laowang) show() {
	logrus.Info("赤裸裸的老王。。。")
}
func (*laowang) cost() int {
	return 0
}

// 衣服装饰器
type clothesDecorator struct {
	// 持有一个被装饰对象
	person Person
}

func (*clothesDecorator) cost() int {
	return 0
}

func (*clothesDecorator) show() {
}

// 夹克
type Jacket struct {
	clothesDecorator
}

func (j *Jacket) cost() int {
	return j.person.cost() + 10
}

func (j *Jacket) show() {
	// 执行已有的方法
	j.person.show()
	logrus.Info("穿上夹克的老王。。。")
}

// 帽子
type Hat struct {
	clothesDecorator
}

func (h *Hat) cost() int {
	return h.person.cost() + 5
}
func (h *Hat) show() {
	logrus.Info("戴上帽子的老王。。。")
}
