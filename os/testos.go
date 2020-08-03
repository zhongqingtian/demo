package main

import (
	"fmt"
	"os"
)

func main()  {
	// 预定义变量，保存命令行参数
	fmt.Println(os.Args)

	// 获得host name
	fmt.Println(os.Hostname())
	fmt.Println(os.Getpid())

	// 获得全部环境变量
	env := os.Environ()
	for k,v := range env{
		fmt.Println(k,v)
	}

	// 终止程序
	//os.Exit(1)

	// 获得一条环境变量
	fmt.Println(os.Getenv("PATH"))

	// 获得当前目录
	dir,err :=os.Getwd()
	fmt.Println(dir,err)

	//创建目录
	err = os.Mkdir(dir+"/new_file",0755)
	fmt.Println(err)

	// 创建目录 包括任何必要的上级目录 可以是多级目录
	err = os.MkdirAll(dir+"/new/demo",0755)
	fmt.Println(err)

	// 删除目录
	err = os.Remove(dir+"/new_file")
	fmt.Println(err)
	err = os.RemoveAll(dir + "/new")
	fmt.Println(err)

	// 创建临时目录
	tem_dir := os.TempDir()
	fmt.Println(tem_dir)
}