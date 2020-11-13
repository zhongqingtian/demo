package pattern

import (
	"github.com/sirupsen/logrus"
)

/*适配器模式 (adaptor)
适配器适合用于解决新旧系统（或新旧接口）之间的兼容问题，而不建议在一开始就直接使用

意图:
适配器模式将一个类的接口，转换成客户期望的另一个接口。适配器让原本接口不兼容的类可以合作无间

关键代码:
适配器中持有旧接口对象，并实现新接口
*/

// 新接口--音乐播放器
type MusicPlayer interface {
	play(fileType string, fileName string)
}

// 旧接口---已经实现好的库
type ExistPlayer struct {
}

func (*ExistPlayer) playMp3(fileName string) {
	logrus.Info("play mp3 :", fileName)
}

func (*ExistPlayer) playWma(fileName string) {
	logrus.Info("play wma:", fileName)
}

/*适配器模式很好的践行了面向对象设计原则里的开闭原则（open/closed principle），新增一个接口时也无需修改老接口，只需多加一个适配层即可*/
// 适配器 加一层封装旧结构
type PlayerAdaptor struct {
	// 持有一个旧接口
	existPlayer ExistPlayer
}

// 实现新的接口
func (player *PlayerAdaptor) play(fileType string, fileName string) {
	switch fileType {
	case "mp3":
		player.existPlayer.playMp3(fileName)
	case "wma":
		player.existPlayer.playWma(fileName)
	default:
		logrus.Info("暂时不支持此类型文件播放")
	}
}
