package pattern

import (
	"testing"
)

func TestObserver(t *testing.T) {
	// 测试观察者模式
	/*customerA := &CustomerA{}
	customerB := &CustomerB{}

	office := &NewsOffice{}

	// 模拟客户订阅
	office.addCustomer(customerA)
	office.addCustomer(customerB)
	// 发放新的报纸
	office.notifyAllCustomer()*/

	// 单例模式
	/*t.Log(GoInstance("单例模式1").Name)
	t.Log(GoInstance("单例模式2").Name) */ // 也是打印出单例模式1字段

	// 策略模式
	/*operator := Operator{}
	operator.setStrategy(&add{})
	result := operator.calculate(1, 2)
	t.Log("add:", result)

	operator.setStrategy(&reduce{})
	result = operator.calculate(2, 1)
	t.Log("reduce:", result)*/

	// 装饰者模式
	/*laowang := &laowang{}

	jacket := &Jacket{}
	jacket.person = laowang // 把实现接口方法的对象赋值给接口，达到实现继承原来功能
	jacket.show()

	hat := &Hat{}
	hat.person = jacket
	hat.show()

	t.Log("cost", hat.cost())*/

	// 适配者模式
	/*player := PlayerAdaptor{}
	player.play("mp3", "死了都要爱")
	player.play("wma", "滴滴")
	player.play("mp4", "复仇者联盟")*/

	// 模版方法模式
	// 做西红柿
	/*xihongshi := &XiHongShi{}
	doCook(xihongshi)

	// 做炒鸡蛋
	chaojidan := &ChaoJiDan{}
	doCook(chaojidan)*/

}

// 八 外观模式
func TestFacade(t *testing.T) {
	startBtn := &StartBtn{}
	startBtn.start()
}

// 九 代理模式
func TestProxy(t *testing.T) {

	station := &Station{3}
	proxy := &StationProxy{station}
	station.sell("小华")
	proxy.sell("派大星")
	proxy.sell("小明")
	proxy.sell("小兰")
}

// 责任链模式
func TestChain(t *testing.T) {
	adHandler := &AdHandler{}
	yellowHandler := &YellowHandler{}
	sensitiveHandler := &SensitiveHandler{}
	// 将责任链串起来
	adHandler.handler = yellowHandler
	yellowHandler.handler = sensitiveHandler

	adHandler.Handle("我是正常内容，我是广告，我是涉黄，我是敏感词，我是正常内容")
}

func TestFactory(t *testing.T) {

	/*store := new(GirlFactoryStore)
	// 提供美国工厂
	store.factory = new(AmericanGirlFactory)
	americanFatGirl := store.createGirl("fat")
	americanFatGirl.weight()

	// 提供中国工厂
	store.factory = new(ChineseGirlFactory)
	chineseFatGirl := store.createGirl("fat")
	chineseFatGirl.weight()*/

	store := new(FactoryStore)
	store.factory = new(GiftFactory) // 实现礼物加成工厂
	giftAddition := store.createFactory(1)
	giftAddition.calculate(AddRule{style: 1})
}
