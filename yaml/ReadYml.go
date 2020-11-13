package yaml

// 配置文件中字母要小写，结构体属性首字母要大写
import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Myconf struct {
	Ipport             string   `yaml:"ipport"`
	StartSendTime      string   `yaml:"startSendTime"`
	SendMaxCountPerDay int      `yaml:"sendMaxCountPerDay"`
	Devices            []Device `yaml:"devices"`
	WarnFrequency      int      `yaml:"warnFrequency"`
	SendFrequency      int      `yaml:"sendFrequency"`
	Sendgifts          []int    `yaml:sendgifts`
	StartTime          string   `yaml:startTime`
	EndTime            string   `yaml:endTime`
}
type Device struct {
	DevId string `yaml:"devId"`
	Nodes []Node `yaml:"nodes"`
}
type Node struct {
	PkId     string  `yaml:"pkId"`
	BkId     string  `yaml:"bkId"`
	Index    string  `yaml:"index"`
	MinValue float32 `yaml:"minValue"`
	MaxValue float32 `yaml:"maxValue"`
	DataType string  `yaml:"dataType"`
}

func ReadYaml() {
	data, _ := ioutil.ReadFile("./conf.yaml") // 先读取 yaml文件
	fmt.Println("原始文件:", string(data))
	t := Myconf{}

	//把然后yaml 形式的字符串解析成struct类型
	yaml.Unmarshal(data, &t)
	fmt.Println("初始化转为结构体数据:", t)
	if t.Ipport == "" {
		fmt.Println("配置文件设置错误！")
		return
	}
	d, _ := yaml.Marshal(&t)
	fmt.Println("恢复原文件：", string(d))
}
