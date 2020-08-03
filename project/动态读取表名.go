package project

import (
	"strings"
	"time"
)

type TableName struct {
	Id        string        `json:"id"`
	Psw       string        `json:"psw"`
	StartTime time.Time     `json:"start_time"`
	Long      time.Duration `json:"long"`
}

var TableClolumns = []string{
	"id",
	"psw",
	"start_time",
	"long",
}

// 定义一个函数类型 返回值是接口类型
type TableNameFileScanner func(tableName *TableName) interface{}

// 定义一个map集合存储函数类型 全局变量做缓存
var tableNameScanners = map[string]TableNameFileScanner{}

func ParseTableScanners(files []string) ([]string, []TableNameFileScanner) {
	if len(files) == 0 {
		return TableClolumns, nil
	}
	// 存储要查询的字段
	cols := make([]string, 0)
	// 存储查询结果 赋值字段的地址
	scanners := make([]TableNameFileScanner, 0)
	for _, f := range files {
		lf := strings.ToLower(f)
		if sc, ok := tableNameScanners[lf]; ok { // 判断缓存中有没有，有就直接获取
			cols = append(cols, lf)
			scanners = append(scanners, sc)
			continue
		}

		switch lf {
		case "id":
			cols = append(cols, lf)
			sc := func(d *TableName) interface{} { //上面定义的函数体 功能是实现获取读取字段的地址
				return &d.Id
			}
			scanners = append(scanners, sc) // 把该方法追加到切片
			tableNameScanners[lf] = sc      // 添加到缓存

		case "psw":
			cols = append(cols, lf)
			sc := func(d *TableName) interface{} {
				return &d.Psw
			}
			scanners = append(scanners, sc)
			tableNameScanners[lf] = sc
		case "start_time":
			cols = append(cols, lf)
			sc := func(d *TableName) interface{} {
				return &d.StartTime
			}
			scanners = append(scanners, sc)
			tableNameScanners[lf] = sc
		case "long":
			cols = append(cols, lf)
			sc := func(d *TableName) interface{} {
				return &d.Long
			}
			scanners = append(scanners, sc)
			tableNameScanners[lf] = sc
		}
	}
	return cols, scanners
}

//_____________
// 传入的 files 包含哪些字段 就表示要 获得查询哪些字段返回地址 如上 &d.StartTime，&d.Long
// 左边是返回 cols 是要操作的字段 []string
// 结果 [id start_time] 0xc0000544e0（地址）
func (r *TableName) Scans(files ...string) ([]string, []interface{}) {
	cols, scanners := ParseTableScanners(files)
	scans := make([]interface{}, len(scanners))
	for i, f := range scanners {
		scans[i] = f(r)
	}
	return cols, scans
}
