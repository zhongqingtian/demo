package main

import (
	"flag"
	"fmt"
	"strings"
)

//定义一个类型，用于增加该类型方法
type sliceValue []string

//new一个存放命令行参数值的slice
func newSliceValue(vals []string, p *[]string) *sliceValue {
	*p = vals
	return (*sliceValue)(p)
}

/*
Value接口：
type Value interface {
    String() string
    Set(string) error
}
实现flag包中的Value接口，将命令行接收到的值用,分隔存到slice里
*/
func (s *sliceValue) Set(val string) error {
	*s = sliceValue(strings.Split(val, ","))
	return nil
}

//flag为slice的默认值default is me,和return返回值没有关系
func (s *sliceValue) String() string {
	*s = sliceValue(strings.Split("default is me", ","))
	return "It's none of my business"
}

var showHelp bool
var k *string
/*
可执行文件名 -slice="java,go"  最后将输出[java,go]
可执行文件名 最后将输出[default is me]
*/
// go run flag.go -config
func main(){
	//os.Args
	var languages []string
	flag.Var(newSliceValue([]string{}, &languages), "slice", "I like programming `languages`") // 给变量设置默认值
	k = flag.String("config","aaa","this is a config")
	flag.Parse() // 读取 命令参数 调用对应 值

	showHelp := flag.Bool("show",true,"show help") // 返回就是 value的指针类型

	if showHelp != nil && *showHelp {
		fmt.Println("出来了吗")
		flag.Usage()
	}



	//打印结果slice接收到的值
	fmt.Println(languages)
}
