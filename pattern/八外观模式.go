package pattern

import "github.com/sirupsen/logrus"

/*意图:
提供了一个统一的接口，用来访问子系统中的一群接口。外观定义了一个高层接口，让子系统更容易使用。

关键代码:
外观层中依次调用子系统的接口

应用实例:

电脑开机时，点击开机按钮，但同时启动了 CPU，内存，硬盘等
用 Java 开发我们经常使用三层结构：controller 控制器层，service 服务层，dao 数据访问层
*/

type CPU struct {
}

func (CPU) start() {
	logrus.Info("启动CPU。。。")
}

// 内存
type Memory struct {
}

func (Memory) start() {
	logrus.Info("启动内存管理。。。")
}

// 硬盘
type Disk struct {
}

func (Disk) start() {
	logrus.Info("启动硬盘。。。")
}

// 开机键
type StartBtn struct {
}

func (StartBtn) start() {
	cpu := &CPU{}
	cpu.start()
	memory := &Memory{}
	memory.start()
	disk := &Disk{}
	disk.start()
}
