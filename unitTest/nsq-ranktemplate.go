package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

// 主函数
func Ranktemplate() {
	strIP1 := "10.10.62.33:9201"
	producer := InitProducer1(strIP1) //根据Ip1地址产生生产者

	/*data, err := proto.Marshal(sendGiftData)
	if err != nil {
		fmt.Println(err)
	}*/
	//推送 默认 “test”的topic 和消息内容
	data := ""
	err := Publish1(producer, "sendgift", string(data))
	if err != nil {
		fmt.Println(err)
	}
	//关闭
	producer.Stop()
}

func SendNotGift() {
	strIP1 := "10.10.62.33:9201"
	producer := InitProducer1(strIP1) //根据Ip1地址产生生产者

	//data, err := proto.Marshal(sendData)
	/*if err != nil {
		fmt.Println(err)
	}*/
	data := ""
	//推送 默认 “test”的topic 和消息内容
	err := Publish1(producer, "statisticsnotgift", string(data))
	if err != nil {
		fmt.Println(err)
	}
	//关闭
	producer.Stop()
}
