package project

import (
	"strings"
)

const (
	RoomRiskColRoomID         = "room_id"
	RoomRiskColEventTotal     = "event_total"
	RoomRiskColIgnoreTotal    = "ignore_total"
	RoomRiskColWarningTotal   = "warning_total"
	RoomRiskColForbiddenTotal = "forbidden_total"
)

// 所有数据列
var RoomRiskColumns = []string{
	RoomRiskColRoomID,
	RoomRiskColEventTotal,
	RoomRiskColIgnoreTotal,
	RoomRiskColWarningTotal,
	RoomRiskColForbiddenTotal,
}

// 表名常量
const (
	TableRoomRisk = "t_roominfo"
)

// 表数据结构
type RoomInfo struct {
	RoomID       string `json:"room_id"`
	EventTotal   uint32 `json:"event_total"`
	IgnoreTotal  uint32 `json:"ignore_total"`
	WarningTotal uint32 `json:"warning_total"`
}

// 接口层 多个表操作时候，可复用很多代码

type Updater interface {
	TableName() string
	Updates(dbOldToNew Updater, fields []string) ([]string, []interface{})
}

// 实现接口
func (*RoomInfo) TableName() string {
	return TableRoomRisk
}

// 创建插入指定某些字段 e是插入的值 fields是要插入的字段片
func (e *RoomInfo) Creates(fields []string) ([]string, []interface{}) {
	cols := make([]string, 0)        //插入字段
	params := make([]interface{}, 0) // 该字段值
	for _, f := range fields {
		lf := strings.ToLower(f)
		switch lf {
		case RoomRiskColRoomID:
			cols = append(cols, lf)
			params = append(params, e.RoomID)
		case RoomRiskColEventTotal:
			cols = append(cols, lf)
			params = append(params, e.EventTotal)
		case RoomRiskColIgnoreTotal:
			cols = append(cols, lf)
			params = append(params, e.IgnoreTotal)
		case RoomRiskColWarningTotal:
			cols = append(cols, lf)
			params = append(params, e.WarningTotal)
		}
	}
	// 最终结果返回要insert的字段组 和 参数值
	return cols, params
}

// e 是要更新的数据          dbOldToNew 数据库查询处出来的旧数据
func (e *RoomInfo) Updates(dbOldToNew Updater, fields []string) ([]string, []interface{}) {
	dbRoomInfo, ok := dbOldToNew.(*RoomInfo) // 类型判断是否为该类型
	if !ok {
		return nil, nil
	}

	cols := make([]string, 0)        // 存返回更新的 字段
	params := make([]interface{}, 0) // 更新参数
	for _, f := range fields {       // 判断要更新哪些字段
		lf := strings.ToLower(f)
		switch lf {
		case RoomRiskColIgnoreTotal:
			if dbRoomInfo != nil && dbRoomInfo.IgnoreTotal != e.IgnoreTotal {
				cols = append(cols, lf)
				params = append(params, e.IgnoreTotal)
				dbRoomInfo.IgnoreTotal = e.IgnoreTotal
			}
		case RoomRiskColWarningTotal:
			if dbRoomInfo != nil && dbRoomInfo.WarningTotal != e.WarningTotal {
				cols = append(cols, lf)
				params = append(params, e.WarningTotal)
				dbRoomInfo.WarningTotal = e.WarningTotal
			}
		}
	}
	return cols, params
}
