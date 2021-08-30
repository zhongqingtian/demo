//Nsq接收测试
package nsq

import (
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

// 消费者
type ConsumerT struct{}

// 主函数
func Receiver(topic, ch string) {
	// 通过 主题 chancel ip地址 创建消费者
	InitConsumer(topic, ch) // http
	time.Sleep(10 * time.Minute)
}

//处理消息 实现 handle接口里面的方法 处理消息内容 这个方法名和参数不可修改
func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	log.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

//初始化消费者
func InitConsumer(topic string, channel string) {
	// Lookupaddr := "127.0.0.1:4161"
	nsqd := "127.0.0.1:4150"
	cfg := nsq.NewConfig()                         //获取nsq 的配置
	cfg.MaxInFlight = 121                          // 允许并发运行的最大消息数  然后也可以自己修改配置字段值
	cfg.LookupdPollInterval = time.Second          //设置重连时间
	c, err := nsq.NewConsumer(topic, channel, cfg) // 新建一个消费者 ，这个消费者监听 topic ,channel 和 一些配置
	if err != nil {
		panic(err)
	}
	c.SetLogger(nil, 0)        //屏蔽系统日志
	c.AddHandler(&ConsumerT{}) // 添加消费者接口 把实体赋给接口 nsq内部自动监控消息分发

	//建立NSQLookupd连接
	err = c.ConnectToNSQD(nsqd)
	// err = c.ConnectToNSQLookupd(Lookupaddr)
	if err != nil {
		panic(err)
	}

	//建立多个nsqd连接
	// if err := c.ConnectToNSQDs([]str{"127.0.0.1:4150", "127.0.0.1:4152"}); err != nil {
	//  panic(err)
	// }

	// 建立一个nsqd连接
	// if err := c.ConnectToNSQD("127.0.0.1:4150"); err != nil {
	//  panic(err)
	// }
}
