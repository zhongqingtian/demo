package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	GetMbInfo()
}

func GetMbInfo() ([]ResInfo, error) {
	var (
		wg           sync.WaitGroup
		metaMap      = make(map[int64]Meta)
		priceMap     = make(map[int64]int)
		downNumerMap = make(map[int64]int64)
		TagMap       = make(map[int64][]Tag)
		templateMap  = make(map[int64]Template)
	)
	start := time.Now()
	ids := SearchRes()
	wg.Add(1)
	go GetMeta(&wg, ids, metaMap)
	wg.Add(1)
	go GetPriceMap(&wg, ids, priceMap)
	wg.Add(1)
	go GetDownNumber(&wg, ids, downNumerMap)
	wg.Add(1)
	go GetRelationTag(&wg, ids, TagMap)
	wg.Add(1)
	go GetTemplate(&wg, ids, templateMap)
	wg.Wait()
	resList := make([]ResInfo, 0)
	for _, id := range ids {
		resInfo := ResInfo{}
		if meta, ok := metaMap[id]; ok {
			resInfo.Name = meta.Name
		}
		if price, ok := priceMap[id]; ok {
			resInfo.Price = price
		}
		if downNumber, ok := downNumerMap[id]; ok {
			resInfo.DownNumber = downNumber
		}
		if tags, ok := TagMap[id]; ok {
			mtag := make([]string, 0)
			for _, tag := range tags {
				mtag = append(mtag, tag.Name)
			}
			resInfo.Tags = mtag
		}
		if templatem, ok := templateMap[id]; ok {
			resInfo.Tid = templatem.Tid
			resInfo.Cat = templatem.Cat
		}
	}
	fmt.Println(time.Since(start))
	return resList, nil
}

// 前端展示结果
type ResInfo struct {
	Id         int64
	Tid        int64
	Cat        string
	Name       string
	Price      int
	DownNumber int64
	Tags       []string
}

func SearchRes() []int64 {
	// do something
	time.Sleep(50 * time.Millisecond)
	return []int64{1, 2, 3, 4, 5, 6, 7}
}

type Meta struct {
	Name string `json:"name"`
}

// 获取基础信息
func GetMeta(wg *sync.WaitGroup, ids []int64, metaMap map[int64]Meta) error {
	defer func() {
		wg.Done()
		err := recover()
		if err != nil {
			//// todo
		}
	}()
	defer func() {
		err := recover()
		if err != nil {
			//// todo
		}
	}()
	// do something
	time.Sleep(50 * time.Millisecond)
	metaMap[1] = Meta{Name: "zhang"}
	metaMap[2] = Meta{Name: "LI"}
	metaMap[10] = Meta{Name: "chen"}
	return nil
}

// 获取价格
func GetPriceMap(wg *sync.WaitGroup, ids []int64, priceMap map[int64]int) error {
	defer func() {
		wg.Done()
		err := recover()
		if err != nil {
			//// todo
		}
	}()
	// do something
	time.Sleep(50 * time.Millisecond)
	priceMap[1] = 10
	priceMap[2] = 81
	priceMap[5] = 89
	priceMap[3] = 0
	return nil
}

func GetDownNumber(wg *sync.WaitGroup, ids []int64, downMap map[int64]int64) error {
	defer func() {
		wg.Done()
		err := recover()
		if err != nil {
			//// todo
		}
	}()
	// do something
	time.Sleep(20 * time.Millisecond)
	downMap[1] = 100
	downMap[2] = 102
	downMap[3] = 103
	return nil
}

type Tag struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// 查询关联标签
func GetRelationTag(wg *sync.WaitGroup, ids []int64, tagsMap map[int64][]Tag) error {
	defer func() {
		wg.Done()
		err := recover()
		if err != nil {
			//// todo
		}
	}()
	// do something
	time.Sleep(100 * time.Millisecond)
	tagsMap[1] = []Tag{{Id: 101, Name: "商业风"}, {Id: 102, Name: "绿色"}}
	tagsMap[2] = []Tag{{Id: 102, Name: "绿色"}}
	return nil
}

type Template struct {
	Tid int64
	Id  int64
	//......
	Cat string
}

// 查询模板信息
func GetTemplate(wg *sync.WaitGroup, ids []int64, tempMap map[int64]Template) error {
	defer func() {
		wg.Done()
		err := recover()
		if err != nil {
			//// todo
		}
	}()
	// do something
	time.Sleep(20 * time.Millisecond)
	tempMap[1] = Template{
		Tid: 100,
		Id:  1,
		Cat: "diagram",
	}
	tempMap[2] = Template{
		Tid:
		101,
		Id:  2,
		Cat: "custom",
	}
	return nil
}
