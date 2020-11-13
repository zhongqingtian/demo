package options

import "fmt"

type SetConfig struct {
	IpAdr string
	Port  int64
	Name  string
	Count int64
}
type options func(*SetConfig)

func GetOptions1() options {
	return func(config *SetConfig) {
		if config == nil {
			config = &SetConfig{}
		}
		config.IpAdr = "127.0.0.1"
		config.Count += 1
	}
}

func GetOptions2() options {
	return func(config *SetConfig) {
		if config == nil {
			config = &SetConfig{}
		}
		config.Name = "lixi"
		config.Count += 1
	}
}

func GetOptions3() options {
	return func(config *SetConfig) {
		if config == nil {
			config = &SetConfig{}
		}
		config.Port = 80
		config.Count += 1
	}
}

func CollectOptions(opts ...options) {
	var optionConfig SetConfig
	for _, opt := range opts { // 更新匿名函数的值
		opt(&optionConfig)
	}
	fmt.Println(optionConfig.Name)
	fmt.Println(optionConfig.IpAdr)
	fmt.Println(optionConfig.Port)
	fmt.Println(optionConfig.Count)
}
