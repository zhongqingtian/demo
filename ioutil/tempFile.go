package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 创建临时目录
	dir, err := ioutil.TempDir("./ioutil", "Testdir") //后缀未知
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(dir) // 用完删除
	fmt.Printf("%s\n", dir)

	//  创建临时文件
	f, err := ioutil.TempFile(dir, "Test")
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(f.Name()) // 用完删除
	fmt.Printf("%s\n", f.Name())
}
