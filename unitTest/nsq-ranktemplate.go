package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"gitlab.xunlei.cn/xllive/proto/nsqtopics"
	pb_nsq "gitlab.xunlei.cn/xllive/proto/pwnsqcommon"
)

// 主函数
func Ranktemplate() {
	strIP1 := "10.10.62.33:9201"
	producer := InitProducer1(strIP1) //根据Ip1地址产生生产者

	sendGiftData := &pb_nsq.SendGiftEvent{
		SendUserid:   649896412,
		AcceptUserid: 209829550,
		RoomId:       "1_10031",
		GiftId:       1097,
		GiftName:     "花",
		Num:          20,
		CostNum:      5,
		CostAddnum:   2,
		SendTime:     1576123261,
		GiftCostnum:  4,
	}
	data, err := proto.Marshal(sendGiftData)
	if err != nil {
		fmt.Println(err)
	}
	//推送 默认 “test”的topic 和消息内容
	err = Publish1(producer, "sendgift", string(data))
	if err != nil {
		fmt.Println(err)
	}
	//关闭
	producer.Stop()
}

func SendNotGift() {
	strIP1 := "10.10.62.33:9201"
	producer := InitProducer1(strIP1) //根据Ip1地址产生生产者

	sendData := &nsqtopics.StatisticsNotGift{
		Id:       "675091047",
		Code:     "ReadWeekUserCharm",
		Order:    1,
		SendTime: 1583143515, // 单位是秒
		Score:    100,
	}
	data, err := proto.Marshal(sendData)
	if err != nil {
		fmt.Println(err)
	}
	//推送 默认 “test”的topic 和消息内容
	err = Publish1(producer, "statisticsnotgift", string(data))
	if err != nil {
		fmt.Println(err)
	}
	//关闭
	producer.Stop()
}
