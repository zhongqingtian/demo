package pattern

import "github.com/sirupsen/logrus"

/*意图:
定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新。
关键代码:
被观察者持有了集合存放观察者 (收通知的为观察者)
应用实例:

报纸订阅，报社为被观察者，订阅的人为观察者
MVC 模式，当 model 改变时，View 视图会自动改变，model 为被观察者，View 为观察者
*/

// 报社 --- 客户
type Customer interface {
	update()
}

type CustomerA struct {
}

func (*CustomerA) update() {
	logrus.Info("我是客户A，我收到报纸了")
}

type CustomerB struct {
}

func (*CustomerB) update() {
	logrus.Info("我是客户B，我收到报纸了")
}

// 报社 -- （被观察者）
type NewsOffice struct {
	customers []Customer
}

func (n *NewsOffice) addCustomer(customer Customer) {
	n.customers = append(n.customers, customer)
}

func (n *NewsOffice) newspaperCome() {
	// 通知所有客户
	n.notifyAllCustomer()
}
func (n *NewsOffice) notifyAllCustomer() {
	for _, customer := range n.customers {
		customer.update()
	}
}
