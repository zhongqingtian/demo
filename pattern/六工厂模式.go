package pattern

//
//import "github.com/sirupsen/logrus"
//
///*意图:
//定义一个创建对象的接口，让其子类自己决定实例化哪一个工厂类，工厂模式使其创建过程延迟到子类进行
//主要解决接口选择问题
//
//关键代码:
//返回的实例都实现同一接口
//
//应用实例:
//
//您需要一辆汽车，可以直接从工厂里面提货，而不用去管这辆汽车是怎么做出来的，以及这个汽车里面的具体实现。
//Hibernate 换数据库只需换方言和驱动就可以。
//*/
//
///*1.简单工厂模式*/
//type Girl interface {
//	weight()
//}
//
//// 胖女孩
//type FatGirl struct {
//}
//
//func (FatGirl) weight() {
//	logrus.Info("80kg")
//}
//
//// 瘦女孩
//type ThinGirl struct {
//}
//
//func (ThinGirl) weight() {
//	logrus.Info("50kg")
//}
//
//type GirlFactory struct {
//}
//
//func (*GirlFactory) CreateGirl(like str) Girl {
//	switch like {
//	case "fat":
//		return &FatGirl{}
//	case "thin":
//		return &ThinGirl{}
//	default:
//		return nil
//	}
//}
//
//// 2.抽象工厂模式
//// 中国胖女孩
//type FatGirl2 struct {
//}
//
//func (FatGirl2) weight() {
//	logrus.Info("chinese girl weight: 80kg")
//}
//
//// 瘦女孩
//type ThinGirl2 struct {
//}
//
//func (ThinGirl2) weight() {
//	logrus.Info("chinese girl weight: 50kg")
//}
//
//type Factory interface {
//	CreateGirl(like str) Girl
//}
//
//// 中国工厂
//type ChineseGirlFactory struct {
//}
//
//func (ChineseGirlFactory) CreateGirl(like str) Girl {
//	switch like {
//	case "fat":
//		return &FatGirl2{}
//	case "thin":
//		return &ThinGirl2{}
//	default:
//		return nil
//	}
//}
//
//// 美国胖女孩
//type AmericanFatGirl struct {
//}
//
//func (AmericanFatGirl) weight() {
//	logrus.Info("American weight: 80kg")
//}
//
//// 美国瘦女孩
//type AmericanThainGirl struct {
//}
//
//func (AmericanThainGirl) weight() {
//	logrus.Info("American weight: 50kg")
//}
//
//// 美国工厂
//type AmericanGirlFactory struct {
//}
//
//func (AmericanGirlFactory) CreateGirl(like str) Girl {
//	switch like {
//	case "fat":
//		return &AmericanFatGirl{}
//	case "thin":
//		return &AmericanThainGirl{}
//	default:
//		return nil
//	}
//}
//
//// 工厂提供者
//type GirlFactoryStore struct {
//	factory Factory
//}
//
//func (store *GirlFactoryStore) createGirl(like str) Girl {
//	return store.factory.CreateGirl(like)
//}
