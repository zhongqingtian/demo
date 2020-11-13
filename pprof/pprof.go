package pprof

import (
	"github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
)

var datas []string

//创建写入 cup.out mem.out 文件
func StartPprof() {
	w, _ := os.Create("cpu.out")
	defer w.Close()
	pprof.StartCPUProfile(w)
	defer pprof.StopCPUProfile()

	w2, _ := os.Create("mem.out")
	defer w2.Close()
	defer pprof.WriteHeapProfile(w2)

	Sum(3, 5)

}

func Sum(a, b int) int {
	return a + b
}

func RunServer() {
	go func() {
		for {
			Add("hello")
			logrus.Info("hello word")
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}

func Add(data string) {
	datas = append(datas, data)
}
