//Nsq发送测试
package nsq

import (
	"bufio"
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
)

var producer *nsq.Producer

// 主函数
func Sender() {
	strIP1 := "10.10.62.33:9201"
	strIP2 := "10.10.62.33:9201"
	InitProducer(strIP1) //根据Ip1地址产生生产者

	running := true

	//读取控制台输入
	reader := bufio.NewReader(os.Stdin)
	for running { //循环读取输入
		data, _, _ := reader.ReadLine() //读入一行
		command := string(data)
		if command == "stop" { //直到接到stop ，退出循环发送消息
			running = false
		}

		//推送 默认 “test”的topic 和消息内容
		for err := Publish("test", command); err != nil; err = Publish("test", command) {
			//切换IP重连
			strIP1, strIP2 = strIP2, strIP1
			InitProducer(strIP1)
		}
	}
	//关闭
	producer.Stop()
}

// 初始化生产者
func InitProducer(str string) {
	var err error
	fmt.Println("address: ", str)
	conf := nsq.NewConfig()                    // 创建一个默认配置
	producer, err = nsq.NewProducer(str, conf) //传入IP地址，创建一个生产者
	if err != nil {
		panic(err)
	}
}

//发布消息
// 传参 topic 和 消息内容
func Publish(topic string, message string) error {
	var err error
	if producer != nil {
		if message == "" { //不能发布空串，否则会导致error
			return nil
		}
		// 调用 生产者 推送 主题 和 消息
		err = producer.Publish(topic, []byte(message)) // 发布消息
		return err
	}
	return fmt.Errorf("producer is nil", err)
}
