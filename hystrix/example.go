package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"net/http"
	"time"
)

func init() {
	// 下面的配置就是我们的请求数量大于等于 10 个并且错误率大于等于 20% 的时候就会触发熔断器开关，熔断器打开 500ms
	// 之后会进入半打开的状态，尝试放一部分请求去访问。
	hystrix.ConfigureCommand("tet", hystrix.CommandConfig{
		// 执行 command 的超时时间
		Timeout: 10,

		// 最大并发量
		MaxConcurrentRequests: 100, // 限流qps

		// 一个统计窗口 10 秒内请求数量
		// 达到这个请求数量后才去判断是否要开启熔断
		RequestVolumeThreshold: 10,

		// 熔断器被打开后
		// SleepWindow 的时间就是控制过多久后去尝试服务是否可用了
		// 单位为毫秒
		SleepWindow: 500,

		// 错误百分比
		// 一个统计窗口 10 秒内请求数量
		// 请求数量大于等于 RequestVolumeThreshold 并且错误率到达这个20%百分比后就会启动熔断
		ErrorPercentThreshold: 20,
	})
}

// https://learnku.com/articles/53090
func main() {
	// 客户端使用:
	runFuc := func() error {
		_, err := http.Get("https://www.baidu.com/")
		if err != nil {
			// fmt.Println("get error:%v", err)
			return err
		}
		// fmt.Println(resp)
		return nil
	}
	callBack := func(err error) error {
		// 出错后的错误处理
		fmt.Println("fallback err: ", err)
		return err
	}

	for i := 1; i < 1000; i++ {
		go func() {
			err := hystrix.Do("test", runFuc, callBack) // 同步 它会阻塞等待，直到执行函数结束或者熔断器返回错误
			if err != nil {
				fmt.Println(err)
			}
			errCh := hystrix.Go("test", runFuc, callBack) // 异步
			fmt.Println(<-errCh)
		}()
	}

	time.Sleep(20 * time.Second)
}
