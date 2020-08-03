package reflect

import (
	"github.com/sirupsen/logrus"
	"reflect"
)

type ControllerInterface interface {
	Init(action string, method string)
}

type Controller struct {
	Action string
	Method string
	Tag    string `json:"tag"`
}

// 实现接口方法
func (c *Controller) Init(action string, method string) {
	c.Action = action
	c.Method = method
}

func TReflect() {
	//初始化
	runController := &Controller{
		Action: "Run1",
		Method: "GET",
	}
	//Controller实现了ControllerInterface方法,因此它就实现了ControllerInterface接口
	var i ControllerInterface
	i = runController
	// 得到接口类型的值
	v := reflect.ValueOf(i)
	logrus.Info("value:", v)

	// 得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	t := reflect.TypeOf(i)
	logrus.Info("type:", t)

	// Elem返回类型的元素类型。
	controllerType := t.Elem()
	tag := controllerType.Field(2).Tag //Field(第几个字段,index从0开始)
	logrus.Info("Tag:", tag)
}
