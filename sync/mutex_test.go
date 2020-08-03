package sync

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestTMutex(t *testing.T) {
	TMutex()
}

func TestNewDataFile(t *testing.T) {
	//简单测试下结果
	var dataFile DataFile
	dataFile, _ = NewDataFile("./mutex_2015_1.dat", 10)

	var d = map[int]Data{
		1: []byte("batu_test1"),
		2: []byte("batu_test2"),
		3: []byte("batu_test3"),
	}

	//写入数据
	for i := 1; i < 4; i++ {
		go func(i int) {
			wsn, _ := dataFile.Write(d[i])
			logrus.Info("write i=", i, ",wsn=", wsn, ",success.")
		}(i)
	}

	//读取数据
	for i := 1; i < 4; i++ {
		go func(i int) {
			rsn, d, _ := dataFile.Read()
			logrus.Info("Read i=", i, ",rsn=", rsn, ",data=", string(d), ",success.")
		}(i)
	}

	time.Sleep(10 * time.Second)
}
