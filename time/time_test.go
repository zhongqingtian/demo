package time

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	//返回现在时间 Time 时间类型
	timeNow := time.Now() //2012-10-31 15:50:13.793654 +0000 UTC
	//log.Println("啦啦啦")
	//Time 时间转化为string
	timeString := timeNow.Format("2006-01-02 15:04:05") //2015-06-15 08:52:32
	t.Log(timeString)

	//获取时间戳
	timestamp := time.Now().Unix() //1504079553

	//时间戳转Time 再转 string
	timeNow = time.Unix(timestamp, 0)                  //2017-08-30 16:19:19 +0800 CST
	timeString = timeNow.Format("2006-01-02 15:04:05") //2015-06-15 08:52:32
	t.Log(timeString)
	//string 转 时间戳
	stringTime := "2017-08-30 16:40:41"
	loc, _ := time.LoadLocation("Local")
	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)
	t.Log(the_time)
	if err == nil {
		unix_time := the_time.Unix() //1504082441
		t.Log(unix_time)
	}
}

func TestGetZT(t *testing.T) {
}

func TestNewTimer(t *testing.T) {
	t.Log(NewTimer())
}

func TestOptimizeTimer(t *testing.T) {
	OptimizeTimer()
}

func TestAfterFunc(t *testing.T) {
	AfterFunc()
}

func TestNewTicker(t *testing.T) {
	NewTicker()
}

func TestStopTicker(t *testing.T) {
	StopTicker()
}

func TestSundate(t *testing.T) {
	Sundate()
}
