package main

import (
	"fmt"
	"time"
)

func main() {
	GetMbInfo()
}

func GetMbInfo() ([]ResInfo, error) {
	start := time.Now()
	ids := SearchRes()
	metaMap, err := GetMeta(ids)
	if err != nil {
		return nil, err
	}
	priceMap, err := GetPriceMap(ids)
	if err != nil {
		return nil, err
	}
	downNumerMap, err := GetDownNumber(ids)
	if err != nil {
		return nil, err
	}
	TagMap, err := GetRelationTag(ids)
	if err != nil {
		return nil, err
	}
	templateMap, err := GetTemplate(ids)
	if err != nil {
		return nil, err
	}
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
func GetMeta(ids []int64) (map[int64]Meta, error) {
	// do something
	time.Sleep(50 * time.Millisecond)
	return map[int64]Meta{1: {Name: "zhang"}, 2: {Name: "LI"}, 3: {Name: "chen"}}, nil
}

// 获取价格
func GetPriceMap(ids []int64) (map[int64]int, error) {
	// do something 
	time.Sleep(50 * time.Millisecond)
	return map[int64]int{1: 10, 2: 1, 3: 3, 4: 1}, nil
}

func GetDownNumber(ids []int64) (map[int64]int64, error) {
	// do something
	time.Sleep(20 * time.Millisecond)
	return map[int64]int64{1: 1000, 2: 200, 3: 26, 6: 410, 7: 102}, nil
}

type Tag struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// 查询关联标签
func GetRelationTag(ids []int64) (map[int64][]Tag, error) {
	// do something
	time.Sleep(100 * time.Millisecond)
	return map[int64][]Tag{1: {{Id: 101, Name: "商业风"}, {Id: 102, Name: "绿色"}, 2: {Id: 2, Name: "简约"}, 3: {Id: 101, Name: "商业风"}, {Id: 102, Name: "绿色"}}}, nil
}

type Template struct {
	Tid int64
	Id  int64
	//......
	Cat string
}

// 查询模板信息
func GetTemplate(ids []int64) (map[int64]Template, error) {
	// do something
	time.Sleep(20 * time.Millisecond)
	return map[int64]Template{1: {
		Tid: 100,
		Id:  1,
		Cat: "diagram",
	}, 2: {
		Tid: 101,
		Id:  2,
		Cat: "custom",
	}}, nil
}
