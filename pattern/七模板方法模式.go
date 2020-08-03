package pattern

import "github.com/sirupsen/logrus"

/*意图:
定义一个操作中的算法的骨架，而将一些步骤延迟到子类中。模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。

关键代码:
通用步骤在抽象类中实现，变化的步骤在具体的子类中实现

应用实例:

做饭，打开煤气，开火，（做饭）， 关火，关闭煤气。除了做饭其他步骤都是相同的，抽到抽象类中实现
spring 中对 Hibernate 的支持，将一些已经定好的方法封装起来，比如开启事务、获取 Session、关闭 Session
*/

type Cooker interface {
	open()
	fire()
	cooke()
	outfire()
	close()
}

// 类似于一个抽象类
type CookMenu struct {
}

func (CookMenu) open() {
	logrus.Info("打开开关")
}

func (CookMenu) fire() {
	logrus.Info("开火")
}

// 做菜，交给具体的子类实现
func (CookMenu) cooke() {
}

func (CookMenu) outfire() {
	logrus.Info("关火")
}

func (CookMenu) close() {
	logrus.Info("关闭开关")
}

// 封装具体步骤
func doCook(cook Cooker) {
	cook.open()
	cook.fire()
	cook.cooke()
	cook.outfire()
	cook.close()
}

type XiHongShi struct {
	CookMenu // 相当于继承了父类
}

func (*XiHongShi) cooke() {
	logrus.Info("做西红柿")
}

type ChaoJiDan struct {
	CookMenu
}

func (ChaoJiDan) cooke() {
	logrus.Info("做炒鸡蛋")
}
