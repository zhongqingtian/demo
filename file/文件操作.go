package main

import (
	"fmt"
	"os"
)

var BasePath string = ""

//写文件代码
func writefile() {
	userFile := "astaxie.txt" //获得文件名称字符串
	b, _ := FileChecker(userFile)
	if !b {
		fout, err := os.Create(userFile) //使用os包下的创建文件夹
		if err != nil {
			fmt.Println(userFile, err)
			return
		}
		defer fout.Close()
		for i := 0; i < 1; i++ {
			fout.WriteString("Just a test!\r\n") //写入
			fout.Write([]byte("Just a test!\r\n"))
		}
	} else {
		//fout, err := os.Open(userFile)
		fout, err := os.OpenFile(userFile, os.O_WRONLY|os.O_APPEND, 0666) //给权限可读可写
		if err != nil {
			fmt.Println(userFile, err)
			return
		}
		defer fout.Close()
		//ret, _ := fout.Seek(0, io.SeekEnd) //查到文件末尾的偏移量
		for i := 0; i < 2; i++ {
			_, er := fout.Write([]byte("我是被追加到追后的\n")) //则从偏移量开始写入
			fmt.Println(er)
		}
	}
}

// 判断资源包是否存在
func FileChecker(filename string) (bool, string) {
	file_path := BasePath + filename
	_, err := os.Stat(file_path)
	if err == nil {
		return true, file_path
	} else {
		return false, "FileChecker:::NotFound " + file_path
	}
}

//读文件
func readfile() {
	userFile := "./astaxie.txt"
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n { //直接读取完内容才推出
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

//读文件
func readfile1() {
	file, err := os.Open("./file.txt")
	//判断文件打开是否打开成功
	if err != nil {
		fmt.Printf("open ./file.txt err : %v\n", err)
	}
	if file != nil {
		defer func(file *os.File) { file.Close() }(file)
	}
	// 创建 b1存储
	var b1 = make([]byte, 102)
	space1, err := file.Read(b1)
	if err != nil {
		fmt.Printf("file read err : %v\n", err)
	}
	fmt.Printf("file read success , 读取 %d 字节。\n", space1)
	fmt.Printf("读取内容：\n%s\n", string(b1))

	b2 := make([]byte, 205)
	space2, err := file.ReadAt(b2, int64(space1))
	if err != nil {
		fmt.Printf("file readat err : %v\n", err)
	}
	fmt.Printf("file readat success , 读取 %d 字节。\n", space2)
	fmt.Printf("读取内容：\n%s\n", string(b2))
}
func main() {
	//writefile()
	//readfile()
	//readfile1()
}
