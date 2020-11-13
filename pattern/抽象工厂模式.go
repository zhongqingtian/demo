package pattern

import (
	"fmt"
	"reflect"
	"strings"
)

/*
在工厂方法模式中，我们通过一个工厂对象来创建一个产品族，具体创建哪个产品，则通过swtich-case的方式去判断。这也意味着该产品组上，每新增一类产品对象，都必须修改原来工厂对象的代码；而且随着产品的不断增多，工厂对象的职责也越来越重，违反了单一职责原则。
抽象工厂模式通过给工厂类新增一个抽象层解决了该问题，如上图所示，FactoryA和FactoryB都实现·抽象工厂接口，分别用于创建ProductA和ProductB。如果后续新增了ProductC，只需新增一个FactoryC即可，无需修改原有的代码；因为每个工厂只负责创建一个产品，因此也遵循了单一职责原则。
*/
// 插件抽象接口定义
type Plugin interface{}

type PluginType string
type InputType Plugin
type FilterType Plugin
type OutputType Plugin

// 输入插件，用于接收消息
type Input interface {
	Plugin
	Receive() string
}

// 过滤插件，用于处理消息
type Filter interface {
	Plugin
	Process(msg string) string
}

// 输出插件，用于发送消息
type Output interface {
	Plugin
	Send(msg string)
}

/*------------------------*/
// 消息管道的定义
type Pipeline struct {
	input  Input
	filter Filter
	output Output
}

// 一个消息的处理流程为 input -> filter -> output
func (p *Pipeline) Exec() {
	msg := p.input.Receive()
	msg = p.filter.Process(msg)
	p.output.Send(msg)
}

/*----定义input、filter、output三类插件接口的具体实现-----*/
// input插件名称与类型的映射关系，主要用于通过反射创建input对象
var inputNames = make(map[string]reflect.Type)

// Hello input插件，接收“Hello World”消息
type HelloInput struct{}

func (h *HelloInput) Receive() string {
	return "Hello World"
}

// 初始化input插件映射关系表
func init() {
	inputNames["hello"] = reflect.TypeOf(HelloInput{})
}

/*----------------*/

// filter插件名称与类型的映射关系，主要用于通过反射创建filter对象
var filterNames = make(map[string]reflect.Type)

// Upper filter插件，将消息全部字母转成大写
type UpperFilter struct{}

func (u *UpperFilter) Process(msg string) string {
	return strings.ToUpper(msg)
}

// 初始化filter插件映射关系表
func init() {
	filterNames["upper"] = reflect.TypeOf(UpperFilter{})
}

/*-----------------*/

// output插件名称与类型的映射关系，主要用于通过反射创建output对象
var outputNames = make(map[string]reflect.Type)

// Console output插件，将消息输出到控制台上
type ConsoleOutput struct{}

func (c *ConsoleOutput) Send(msg string) {
	fmt.Println(msg)
}

// 初始化output插件映射关系表
func init() {
	outputNames["console"] = reflect.TypeOf(ConsoleOutput{})
}

/*-----定义插件抽象工厂接口，以及对应插件的工厂实现--------*/

type Config struct {
	Name string
}

// 插件抽象工厂接口
type Factory2 interface {
	Create(conf Config) Plugin
}

// input插件工厂对象，实现Factory接口
type InputFactory struct{}

// 读取配置，通过反射机制进行对象实例化
func (i *InputFactory) Create(conf Config) Plugin {
	t, _ := inputNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

// filter和output插件工厂实现类似
type FilterFactory struct{}

func (f *FilterFactory) Create(conf Config) Plugin {
	t, _ := filterNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

type OutputFactory struct{}

func (o *OutputFactory) Create(conf Config) Plugin {
	t, _ := outputNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}

// 保存用于创建Plugin的工厂实例，其中map的key为插件类型，value为抽象工厂接口
var pluginFactories = make(map[PluginType]Factory)

// 根据plugin.Type返回对应Plugin类型的工厂实例
func factoryOf(t PluginType) Factory {
	factory, _ := pluginFactories[t]
	return factory
}

// pipeline工厂方法，根据配置创建一个Pipeline实例
/*func Of(conf Config) *Pipeline {
	p := &Pipeline{}
	p.input = factoryOf(InputType).Create(conf.Input).(plugin.Input)
	p.filter = factoryOf(plugin.FilterType).Create(conf.Filter).(plugin.Filter)
	p.output = factoryOf(plugin.OutputType).Create(conf.Output).(plugin.Output)
	return p
}

// 初始化插件工厂对象
func init() {
	pluginFactories[plugin.InputType] = &plugin.InputFactory{}
	pluginFactories[plugin.FilterType] = &plugin.FilterFactory{}
	pluginFactories[plugin.OutputType] = &plugin.OutputFactory{}
}*/
