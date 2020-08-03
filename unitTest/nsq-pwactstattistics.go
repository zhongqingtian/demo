package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	pbAuPayBag "gitlab.xunlei.cn/xllive/proto/aupaybag"
	"gitlab.xunlei.cn/xllive/proto/nsqtopics"
	"gitlab.xunlei.cn/xllive/proto/pwnsqcommon"
	pb_nsq "gitlab.xunlei.cn/xllive/proto/pwnsqroom"
	"time"
)

// 进房间
func InRoom() {
	strIP1 := "10.10.62.33:9201"
	producer := InitProducer1(strIP1) //根据Ip1地址产生生产者
	inRoom := &pb_nsq.InRoomEvent{
		Roomid:     "10011",
		Message:    "",
		Userid:     "655223076",
		Timestamp:  1572938105,
		UniqueId:   "",
		DeviceId:   "",
		DeviceType: "",
	}
	data, err := proto.Marshal(inRoom)
	if err != nil {
		fmt.Println(err)
	}
	//推送 默认 “test”的topic 和消息内容
	err = Publish1(producer, "inroom", string(data))
	if err != nil {
		fmt.Println(err)
	}
	//关闭
	producer.Stop()
}

// 离开房间
func OutRoom() {
	strIP1 := "10.10.62.33:9201"
	producer := InitProducer1(strIP1) //根据Ip1地址产生生产者
	outRoom := &pb_nsq.OutRoomEvent{
		Roomid:    "10011",
		Message:   "",
		Userid:    "655223076",
		Timestamp: 1572938205,
	}
	data, err := proto.Marshal(outRoom)
	if err != nil {
		fmt.Println(err)
	}
	//推送 默认 “test”的topic 和消息内容
	err = Publish1(producer, "outroom", string(data))
	if err != nil {
		fmt.Println(err)
	}
	//关闭
	producer.Stop()
}

func Sendchat() {
	strIP1 := "10.10.62.33:9201"
	producer := InitProducer1(strIP1) //根据Ip1地址产生生产者
	sendChat := &xllive_pwnsqcommon.PwSendChatEvent{
		Userid:    655223076,
		RoomId:    "10011",
		Content:   "bilibili",
		Timestamp: 1572938205,
	}
	data, err := proto.Marshal(sendChat)
	if err != nil {
		fmt.Println(err)
	}

	err = Publish1(producer, "pwsendchat", string(data))
	if err != nil {
		fmt.Println(err)
	}
	//关闭
	producer.Stop()
}

// 70811701 68574493 71067961 71068012 1_50095 1_50089
func SendGift(SendUserid uint64, AcceptUserid uint64, roomId string, GiftId uint64, Num uint64, time uint64) {
	strIP1 := "10.10.62.33:9201"
	producer := InitProducer1(strIP1) //根据Ip1地址产生生产者
	sendgift := &xllive_pwnsqcommon.SendGiftEvent{
		SendUserid:   SendUserid,
		AcceptUserid: AcceptUserid,
		RoomId:       "1_50095",
		GiftId:       GiftId,
		GiftName:     "花",
		Num:          Num,
		CostNum:      5,
		CostAddnum:   2,
		SendTime:     time,
		GiftCostnum:  10,
	}
	data, err := proto.Marshal(sendgift)
	if err != nil {
		fmt.Println(err)
	}

	err = Publish1(producer, "sendgift", string(data))
	if err != nil {
		fmt.Println(err)
	}
	//关闭
	producer.Stop()
}

func NotGift() {
	strIP1 := "10.10.62.33:9201"
	producer := InitProducer1(strIP1) //根据Ip1地址产生生产者
	sendChat := &nsqtopics.StatisticsNotGift{
		Id:       "596694974",
		Code:     "61archBinRank",
		Order:    0,
		SendTime: uint64(time.Now().Unix()),
		Score:    100,
	}
	data, err := proto.Marshal(sendChat)
	if err != nil {
		fmt.Println(err)
	}

	err = Publish1(producer, "statisticsnotgift", string(data))
	if err != nil {
		fmt.Println(err)
	}
	//关闭
	producer.Stop()
}

func Pay() {
	strIP1 := "10.10.62.35:9201"
	producer := InitProducer1(strIP1) //根据Ip1地址产生生产者
	sendChat := &pbAuPayBag.PayEvent{
		Orderid:  "12",
		Userid:   "160222585",
		Money:    "1",
		AddTime:  "2020-10-12 20:02:03",
		PageFrom: "andrin",
		Uniqid:   "11",
	}
	data, err := json.Marshal(sendChat)
	if err != nil {
		fmt.Println(err)
	}

	err = Publish1(producer, "pay", string(data))
	if err != nil {
		fmt.Println(err)
	}
	//关闭
	producer.Stop()
}
