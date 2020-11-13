package promethues

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"log"
	"math"
	"net/http"
	"time"
)

var (
	MyTestCounter = prometheus.NewCounter(prometheus.CounterOpts{
		//因为Name不可以重复，所以建议规则为："部门名_业务名_模块名_标量名_类型"
		Name: "my_test_counter", //唯一id，不可重复Register()，可以Unregister()
		Help: "my test counter", //对此Counter的描述
	})
	MyTestGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "my_test_gauge",
		Help: "my test gauge",
	})
	MyTestHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "my_test_histogram",
		Help:    "my test histogram",
		Buckets: prometheus.LinearBuckets(20, 5, 5), //第一个桶20起，每个桶间隔5，共5个桶。 所以20, 25, 30, 35, 40
	})
	MyTestSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "my_test_summary",
		Help:       "my test summary",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}, //返回五分数， 九分数， 九九分数
	})
)

const (
	statusAPIError = 422
)

func PromRun() {

	//不能注册多次Name相同的Metrics
	//MustRegister注册失败将直接panic()，如果想捕获error，建议使用Register()
	prometheus.MustRegister(MyTestCounter)
	prometheus.MustRegister(MyTestGauge)
	prometheus.MustRegister(MyTestHistogram)
	prometheus.MustRegister(MyTestSummary)

	go func() {
		var i float64
		for {
			i++
			MyTestCounter.Add(10000)                                                  //每次加常量
			MyTestGauge.Add(i)                                                        //每次加增量
			MyTestHistogram.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10) //每次观察一个18 - 42的量
			MyTestSummary.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)

			time.Sleep(time.Second)
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil)) //多个进程不可监听同一个端口
}

type PromeClient struct {
	apiClient v1.API
}

func NewApiClient() *PromeClient {
	httpClient, err := api.NewClient(api.Config{
		Address:      "http://promsre.cn", // 不用指定端口，默认就行
		RoundTripper: api.DefaultRoundTripper,
	})
	if err != nil {
		logrus.Warn(err)
	}
	promAPI := v1.NewAPI(httpClient) // 创建client
	return &PromeClient{apiClient: promAPI}
}

/*var promeClient *PromeClient

func init() {
	promeClient = NewApiClient()
}*/

func (p *PromeClient) Series(str ...string) {
	ctx := context.Background()
	start := time.Unix(1604887810, 0)
	end := start.Add(time.Minute / 2)
	res, _, err := p.apiClient.Series(ctx, str, start, end)
	logrus.Info(err)
	fmt.Println("count=", len(res))
	for i, re := range res {
		fmt.Println(i, re)
	}
}

func (p *PromeClient) Query(str string) {
	ctx := context.Background()
	st := time.Unix(1604887810, 0)
	res, _, err := p.apiClient.Query(ctx, str, st)
	if err != nil {
		logrus.Info(err)
	}
	fmt.Println(res)
}

func Get() {

}
