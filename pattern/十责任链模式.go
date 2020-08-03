package pattern

import (
	"github.com/sirupsen/logrus"
	"strings"
)

/*
意图:
为某个请求创建一个对象链，每个对象依次检查此请求，并对其进行处理，或者将它传给链中的下一个对象

关键代码:
责任链中每个对象都拥有同一个父类（或接口）

应用实例:

消息过滤器，权限拦截器
用户发帖内容进行广告过滤，涉黄过滤，敏感词过滤等
*/

type Handler interface {
	Handle(content string)
	next(handler Handler, content string)
}

// 广告过滤
type AdHandler struct {
	handler Handler
}

func (ad *AdHandler) Handle(content string) {
	logrus.Info("执行广告过滤。。。")
	newContent := strings.Replace(content, "广告", "**", 1)
	logrus.Info(newContent)
	ad.next(ad.handler, newContent)
}

func (ad *AdHandler) next(handler Handler, content string) {
	if ad.handler != nil {
		ad.handler.Handle(content)
	}
}

// 涉黄过滤
type YellowHandler struct {
	handler Handler
}

func (yellow *YellowHandler) Handle(content string) {
	logrus.Info("执行涉黄过滤。。。")
	newContent := strings.Replace(content, "涉黄", "**", 1)
	logrus.Info(newContent)
	yellow.next(yellow.handler, newContent)
}

func (yellow *YellowHandler) next(handler Handler, content string) {
	if yellow.handler != nil {
		yellow.handler.Handle(content)
	}
}

// 敏感词过滤
type SensitiveHandler struct {
	handler Handler
}

func (sensitive *SensitiveHandler) Handle(content string) {
	logrus.Info("执行敏感词过滤。。。")
	newContent := strings.Replace(content, "敏感词", "***", 1)
	logrus.Info(newContent)
	sensitive.next(sensitive.handler, newContent)
}

func (sensitive *SensitiveHandler) next(handler Handler, content string) {
	if sensitive.handler != nil {
		sensitive.handler.Handle(content)
	}
}
