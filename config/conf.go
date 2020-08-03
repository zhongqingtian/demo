package config

import (
	"github.com/lxmgo/config"
	"github.com/sirupsen/logrus"
)

func LoadCfg() {
	config, err := config.NewConfig("./conf.ini")
	if err != nil {
		logrus.Info(err)
	}
	result, err := config.Int("port") // result is int 8080
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info("port = ", result)

	str1 := config.String("host")
	// result is string "act.wiki"
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info("host = ", str1)
	str2 := config.String("mysql::host") // 读取指定匹配的规则
	// result is string "127.0.0.1"
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info("mysql.host = ", str2)
}
