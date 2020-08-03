package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Items struct {
	Confidence    float64 `json:"confidence"`//自信度
	Positive_prob float64 `json:"positive_prob"`//积极概率
	Negative_prob float64 `json:"negative_prob"`//消极概率
	Sentiment  int `json:"sentiment"`// 0表示负面 1 中性 3 积极
}
type Emotion struct {
	Log_id string `json:"log_id"`//请求唯一标识码
	Text  string `json:"text"`//原文本
	Items []Items `json:"items"`
}

func httpPost(text string) *Emotion{

    fromstr :=`{ "text": "`+text+`"}`
	jsonStr :=[]byte(fromstr)
	url:= "https://aip.baidubce.com/rpc/2.0/nlp/v1/sentiment_classify?charset=UTF-8&access_token=24.fbe6cac4e237e16c34a834f6b04e4f30.2592000.1569067756.282335-17066930"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	statuscode := resp.StatusCode
	hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	var data  = Emotion{}
	fmt.Println(string(body))
	json.Unmarshal(body,&data)
	fmt.Println(data.Text)
	fmt.Println(statuscode)
	fmt.Println(hea)

	return &data
}

func main()  {
	key :="我要去死了"
	emotion := httpPost(key)
    fmt.Println(emotion.Items[0].Positive_prob)
}