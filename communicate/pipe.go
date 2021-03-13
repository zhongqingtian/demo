package communicate

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
	"time"
)

/*
1、golang实现并发子进程通信(实例)
2、创建一个并发ping程序，然后将ping的执行结果通过管道(pipe)传递给主程序
3、主程序读取结果，然后打印，最后通过主程序发送系统信号(syscall.Signal)来结束ping。
*/
func Ping() {
	cmd1 := exec.Command("ping", "www.baidu.com") //创建子进程
	ppReader, err := cmd1.StdoutPipe() // 创建管道,ppReader为io.ReaderCloser类型
	defer ppReader.Close()
	var bufReader = bufio.NewReader(ppReader) // 创建带缓冲的Reader
	if err != nil {
		fmt.Printf("create cmd stdoutpipe failed,error:%s\n", err)
		os.Exit(1)  // 中断整个程序  属于异常处理
	}
	err = cmd1.Start()
	if err != nil {
		fmt.Printf("cannot start cmd1,error:%s\n", err)
		os.Exit(1)  // 中断整个程序  属于异常处理
	}
	go func() {
		var buffer []byte = make([]byte, 4096)
		for {
			n, err := bufReader.Read(buffer) // 读取数据到buffer中
			if err != nil {
				if err == io.EOF { //管道关闭后会出现io.EOF错误
					fmt.Printf("pipi has Closed\n")
					break
				} else {
					fmt.Println("Read content failed")
				}
			}
			fmt.Print(string(buffer[:n]))
		}
	}()
	time.Sleep(10 * time.Second) //让子进程运行10秒  相当于主程序让出 计算空间给其他进程
	err = stopProcess(cmd1)   //    停止子进程
	if err != nil {
		fmt.Printf("stop child process failed,error:%s", err)
		os.Exit(1)
	}
	cmd1.Wait() //能让cmd1执行到结束,测试stopProcess是否成功,如果没有stopProcess,程序和子进程会一直执行
	time.Sleep(1 * time.Second) //让协程能读取关闭子进程时最后一次传输的数据
}

func stopProcess(cmd *exec.Cmd) error {
	pro, err := os.FindProcess(cmd.Process.Pid) //通过pid获取子进程
	if err != nil {
		return err
	}
	err = pro.Signal(syscall.SIGINT)  // 给子进程发送信号使之结束
	if err != nil {
		return err
	}
	fmt.Printf("结束子进程%s成功\n", cmd.Path)
	return nil
}
