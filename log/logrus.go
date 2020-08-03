package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

func LogOutput() {
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	logrus.Debug("Useful debugging information.")
	logrus.Info("Something noteworthy happened!")
	logrus.Warn("You should probably take a look at this.")
	logrus.Error("Something failed but I'm not quitting.")
	//logrus.Fatal("Bye.") //log之后会调用os.Exit(1)
	//logrus.Panic("I'm bailing.") //log之后会panic()

	//在多个地方使用日志可以创建logger
	log := logrus.New()
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
		log.Infof(fmt.Sprintf("%v fail", err), "出错了")
	}

	/*log.WithFields(logrus.Fields{
		"filename": "123.txt",
	}).Info("打开文件夹")
	log.Info("Failed to log to file, using default stderr")
	log.Infof(fmt.Sprintf("%v fail", err), "出错了")*/

	request_id := "32645"
	user_ip := "admin"
	entry := logrus.WithFields(logrus.Fields{"request_id": request_id, "user_ip": user_ip})
	entry.Info("something happened on that request")
	entry.Warn("something not great happened")
	log.WithFields(logrus.Fields{
		"filename": "123.txt",
	}).Info("打开文件夹")
	log.Info("Failed to log to file, using default stderr")
	log.Infof(fmt.Sprintf("%v fail", err), "出错了")
}

func init() {
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	logrus.SetOutput(os.Stdout)
	//设置最低loglevel
	logrus.SetLevel(logrus.InfoLevel)
}
