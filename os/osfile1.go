package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	/*proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}*/
	/*file, err := os.Create("./os/proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, p := range proverbs {
		// file 类型实现了 io.Writer
		n, err := file.Write([]byte(p))  //把字符写入当前文档
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	fmt.Println("file write done")*/

	// 获得当前目录
	dir,err := os.Getwd()
    fmt.Println(dir,err)

	file := dir + "/new.txt"
	var fh  *os.File

	fi,_ := os.Stat(file) // 读取文件信息，不是文件夹
	fmt.Println("文件名：",fi.Name())
	fmt.Println("文件大小：",fi.Size())
	fmt.Println("文件修改权限：",fi.Mode())
	fmt.Println("文件是否目录：",fi.IsDir())
	if fi==nil {
		fh,_ = os.Create(file) // 文件不存在就创建
	}else {
		fh,_ = os.OpenFile(file,os.O_RDWR,0666) // 文件存在就打开
	}

	w := []byte("hello go language"+time.Now().String())
	n,err := fh.Write(w)
	fmt.Println(n,err)

	// 设置下次读写位置
	ret,err := fh.Seek(0,0)
	fmt.Println("当前文件指针位置",ret,err)

	b := make([]byte,1024)
	n,err = fh.Read(b)
	fmt.Println(n,err,string(b))

	fh.Close()

}
