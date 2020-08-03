package pattern

import (
	"github.com/sirupsen/logrus"
	"time"
)

type AddRule struct {
	code     []string // 必定指定某个code，或者多个活动code
	style    uint32   // 选择哪种加成工厂模式
	scale    float32  // 比例
	addPoint float32  // 直接加分
	roomId   string
	userId   string
	startAt  time.Time
	endTime  time.Time
}

// 用于存储缓存
type RuleMaps struct {
	RuleMaps map[string]AddRule `json:"rule_maps"`
}

// 抽象工厂接口
type Factory interface {
	CreateCalculator(style uint32) ExtraPoint
}

// 工厂提供者
type FactoryStore struct { // 送礼加成工厂、非礼物加成工厂 要提前实现工厂接口方法CreateCalculator
	factory Factory
}

func (store *FactoryStore) createFactory(style uint32) ExtraPoint {
	return store.factory.CreateCalculator(style)
}

type ExtraPoint interface { // 实现额外积分添加接口
	calculate(rule AddRule) float64
}

// 礼物添加额外积分工厂
type GiftFactory struct {
}

// 实现生成礼物积分模式计算器
func (GiftFactory) CreateCalculator(style uint32) ExtraPoint { // 某个模式又实现添加积分接口
	switch style {
	case 1:
		return &RoomPoint{} // 每次添加一种加成模式要在这里实现
	case 2:
		return &UserPoint{}
	default:
		return nil
	}
}

// 房间类礼物额外加成
type RoomPoint struct {
}

// 房间类实现了计算积分
func (RoomPoint) calculate(rule AddRule) float64 {
	logrus.Info("RoomPoint实现过程")
	panic("")
}

// 礼物类额外加成
type UserPoint struct {
}

// 房间类实现了计算积分
func (UserPoint) calculate(rule AddRule) float64 {
	logrus.Info("UserPoint实现过程")
	panic("")
}

// 礼物添加额外积分工厂
type OtherRankFactory struct {
}

// 实现生成非礼物礼物积分模式计算器
func (OtherRankFactory) CreateCalculator(style uint32) ExtraPoint { // 某个模式又实现添加积分接口
	switch style {
	case 3:
		return &OtherRoomPoint{}
	case 4:
		return &OtherUserPoint{}
	default:
		return nil
	}
}

// 房间类礼物额外加成
type OtherRoomPoint struct {
}

// 房间类实现了计算积分
func (OtherRoomPoint) calculate(rule AddRule) float64 {
	logrus.Info("OtherRoomPoint实现过程")
	panic("")
}

// 礼物类额外加成
type OtherUserPoint struct {
}

// 房间类实现了计算积分
func (OtherUserPoint) calculate(rule AddRule) float64 {
	logrus.Info("OtherUserPoint实现过程")
	panic("")
}
