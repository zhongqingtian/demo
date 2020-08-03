package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"
)

func Pay() {
	strIP1 := "10.10.62.35:9201"
	producer := InitProducer1(strIP1) //根据Ip1地址产生生产者

	/*data, err := json.Marshal(sendChat)
	if err != nil {
		fmt.Println(err)
	}*/
	data := ""
	err := Publish1(producer, "pay", string(data))
	if err != nil {
		fmt.Println(err)
	}
	//关闭
	producer.Stop()
}
