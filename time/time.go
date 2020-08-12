package time

import (
	"github.com/sirupsen/logrus"
	"time"
)

// 构造一个定时器，到达时间后，会把当前时间输出到chan
func NewTimer() time.Time {
	t := time.NewTimer(3*time.Second + 3*time.Millisecond)
	expire := <-t.C // 这个chan会一直阻塞直到定时器到到期返回当前时间

	return expire
}

func OptimizeTimer() {
	ch1 := make(chan int, 1000)
	sign := make(chan byte, 1)

	//给ch1里面写入数据
	for i := 0; i < 1000; i++ {
		ch1 <- i
	}

	// 单独起一个goroutine执行select

	go func() {
		var e int
		ok := true

		var timer *time.Timer

		for {
			select {
			case e = <-ch1:
				logrus.Infof("ch1 = %d", e)
			case <-func() <-chan time.Time {
				if timer == nil {
					// 初始化1秒的定时器
					timer = time.NewTimer(2 * time.Millisecond)
				} else {
					// 复用，重设定时器
					timer.Reset(2 * time.Millisecond)
				}
				// 等待定时器事件到来，返回结果
				return timer.C
			}():
				logrus.Info("TimeOut.")
				ok = false
				break
			}
			// 终止for循环
			if !ok {
				sign <- 0
				break
			}
		}
	}()

	// 通过sign通道数据，为了等待select的Goroutine执行
	<-sign
}

func AfterFunc() {
	var t *time.Timer

	f := func() {
		logrus.Infof("Expiration time : %v.\n", time.Now())
		logrus.Infof("C`s len: %d\n", len(t.C))
	}

	t = time.AfterFunc(1*time.Second, f) // 以一个新创建一个协程的形式，不会阻塞当前主协程
	//让当前Goroutine 睡眠2s，确保大于内容的完整
	//这样做原因是，time.AfterFunc的调用不会被阻塞。它会以一部分的方式在到期事件来临执行我们自定义函数f。
	time.Sleep(2 * time.Second) // 睡眠阻塞2s等待time.AfterFunc方法协程结果返回打印
}

// 断续器
func NewTicker() {
	// 初始化断续器，时间间隔1s
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)

	go func() {
		for t := range ticker.C { // for 循环 chan 会阻塞等待数据，直到该chan关闭销毁 ,close(ch)
			logrus.Info("Tick at", t)
		}
	}()

	time.Sleep(time.Second * 5) // 阻塞，则执行次数为sleep的时间5s/ticker的时间间隔1s=5次，因为ticker是周期性到来
	ticker.Stop()

	logrus.Info("Ticker stopped")
}

func StopTicker() {
	//初始化断续器,间隔2s
	var ticker *time.Ticker = time.NewTicker(100 * time.Millisecond)
	// num 为指定的执行次数
	num := 2
	c := make(chan int, num)
	go func() {
		for t := range ticker.C {
			c <- 1 //　当chan 满2次会阻塞，等待调用ticker.Stop()退出for循环
			logrus.Info("Tick at", t)
		}
	}()

	time.Sleep(1500 * time.Millisecond)
	ticker.Stop()
	logrus.Info("Ticker stopped")
}

func Sundate() {
	now := time.Now()
	thisSunDay := time.Sunday
	ThisWeekDay := now.Weekday()
	offSet := int(thisSunDay - ThisWeekDay)
	sunDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offSet)
	firstDurTime := sunDate.Sub(now)
	logrus.Warn(firstDurTime)
}

func ParseDate() {
	date := time.Date(2020, 10, 21, 12, 12, 10, 0, time.Local)
	logrus.Print(date.Format("01/02-15:04"))
}
