package main

import (
	"fmt"
	"os"
)
//创建和删除文件
func createAnddelete()  {
	os.Mkdir("./astaxie", 0777)
	os.MkdirAll("./astaxie/test1/test2", 0777)
	err := os.Remove("./astaxie")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("./astaxie")
}

//打开文件，写文件
func openandwrite()  {
	userFile := "astaxie.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}
}

func main(){
   createAnddelete()
}
