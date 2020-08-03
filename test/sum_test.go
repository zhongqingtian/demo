package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type A struct {
	AA int `json:"aa"`
}


var ch = make(chan int,3)
var wg sync.WaitGroup

type MyType time.Duration

func (mt MyType) String() string {
	return fmt.Sprintf("时间消耗:%v ms", int64(time.Duration(mt) / time.Millisecond))
}

func TestSum(t *testing.T)  {
	//start := time.Now().UnixNano()
	start := time.Now()
	wg.Add(1)
	go sum([]int{4,5,9,4,51},ch)
	wg.Add(1)
	go sum([]int{4,5,9,4,51},ch)
	wg.Add(1)
	go sum([]int{4,5,9,4,51},ch)
	wg.Wait()
	close(ch) //用完关闭通道
	time.Sleep(50 *time.Millisecond)
	end := time.Now()

	interval := end.Sub(start)
	myInterval := MyType(interval)

	fmt.Println(myInterval)
	//time.Sleep(time.Second)
	getChanNum(ch)
	//time.Sleep(time.Second)
	//stop := make(chan bool)

	//go printOUt(ch,stop)

	//time.Sleep(10)
	//close(stop)
	//time.Sleep(time.Second)

}



func sum(s []int, ch chan int)  { //求和
	//wg.Add(1)
	sum := 0
	for _, v := range s{
		sum += v
	}

	for i:=0;i<10000;i++ {

	}
	ch <- sum //最后把结果推到通道
	wg.Done()   // 等同 wg.Add(-1)
}

func getChanNum(c chan int)  {
	fmt.Println(len(c))

	for k:= range c {
		fmt.Println(k)
	}

}

func printOUt(c chan int,stop chan bool) {
	ticker := time.NewTicker(2*time.Second)
	for {
		select {
		case <- ticker.C:
			//
		//	print(time.Now())
		case retValue := <- c:
			// do
			print(retValue)
		case <- stop:
			print("receive stop sing")
			return
		}
	}
}