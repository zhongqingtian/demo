package main

import (
	"demo/gobreaker"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	// "github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker

func init() {
	st := gobreaker.Settings{
		Name:          "",               // 熔断器名称
		MaxRequests:   10,               // 最大请求数 （半开启状态会限流）
		Interval:      1 * time.Second, // 统计周期
		Timeout:       1 * time.Second,  // 进入熔断后的超时时间
		ReadyToTrip:   nil,              // 通过Counts 判断是否开启熔断。需要自定义
		OnStateChange: nil,              // 状态修改时的钩子函数
		IsSuccessful:  nil,              // 可自定义判断
	}
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 10 && failureRatio >= 0.1
	}

	cb = gobreaker.NewCircuitBreaker(st)
}

// Get wraps http.Get in CircuitBreaker.
func Get(url string) ([]byte, error) {
	f := func() (interface{}, error) { // 业务处理
		// return nil,fmt.Errorf("tt")
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		rand.Seed(time.Now().UnixNano())
		kk := rand.Intn(10000)
		ratia := kk%2
		// fmt.Println(ratia)
		if  ratia == 1 {
			// fmt.Println("err 111")
			return nil, fmt.Errorf("人为出错")
		}
		return body, nil
	}
	body, err := cb.Execute(f)
	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}

func main() {
	for i := 1; i < 1000; i++ {
		go func(i int) {
			_, err := Get("http://www.google.com/robots.txt")
			if err != nil {
				fmt.Println(err)
			}
			// fmt.Println("i=",i)
		}(i)
	}
 time.Sleep(10*time.Second)
	// fmt.Println(string(body))
}
