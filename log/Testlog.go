package log

import (
	"fmt"
	"log"
	"os"
)

func TestLog() {
	fmt.Println("begin TestLog ...")

	file, err := os.Create("test.log") // 创建日志文件
	if err != nil {
		// 打印日志并且退出
		log.Fatalln("fail to create test.log file!")
	}

	// 创建logger 对象 这种方式会显示触发日志文件行数
	logger := log.New(file, "我是个日志前缀", log.LstdFlags|log.Llongfile)
	log.Println(" log.Println 1.Println log with log.LstdFlags ...")
	logger.Println("logger.Println 1.Println log with log.LstdFlags ...")

	logger.SetFlags(log.LstdFlags) // 设置每行日志打印格式 仅显示时间

	log.Println("log.Println  2.Println log without log.LstdFlags ...")
	logger.Println("logger.Println  2.Println log without log.LstdFlags ...")

	log.Println("log  4.Println log without log.LstdFlags ...")
	logger.Println("logger  4.Println log without log.LstdFlags ...")

	log.Fatal("5.std Fatal log without log.LstdFlags ...") // 下面的不会打印，此行会终止程序
	fmt.Println("5 Will this statement be execute ?")
	logger.Fatal("5.Fatal log without log.LstdFlags ...")

	// 并发测试   log包是并发goroutine安全的，而fmt不是
	log.Printf(" %s", "log print 11111111111111111")
	log.Printf(" %s", "log print 22222222222222222")

	fmt.Println("fmt print 11111111111")
	fmt.Println("fmt print 22222222222")
}
