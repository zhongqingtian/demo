package pattern

import "github.com/sirupsen/logrus"

/*意图:
代理模式为另一个对象提供一个替身或者占位符以控制对这个对象的访问

关键代码:
代理类和被代理类实现同一接口，代理类中持有被代理类对象

应用实例:

火车票的代理售票点。代售点就是代理，它拥有被代理对象的部分功能 — 售票功能
明星的经纪人，经纪人就是代理，负责为明星处理一些事务。
*/

type Seller interface {
	sell(name string)
}

// 火车站
type Station struct {
	stock int //库存
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		logrus.Infof("代理点中：%s买了一张票,剩余：%d \n", name, station.stock)
	} else {
		logrus.Info("票已售空")
	}
}

// 火车站代理点
type StationProxy struct {
	station *Station // 持有一个火车站对象
}

func (proxy *StationProxy) sell(name string) {
	if proxy.station.stock > 0 {
		proxy.station.stock--
		logrus.Infof("代理点中：%s买了一张票,剩余：%d \n", name, proxy.station.stock)
	} else {
		logrus.Info("票已售空")
	}
}
