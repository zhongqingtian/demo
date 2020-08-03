package project

import (
	"strings"
	"time"
)

const (
	RoomRiskColRoomID           = "room_id"
	RoomRiskColEventTotal       = "event_total"
	RoomRiskColIgnoreTotal      = "ignore_total"
	RoomRiskColWarningTotal     = "warning_total"
	RoomRiskColForbiddenTotal   = "forbidden_total"
	RoomRiskColForbiddenStartAt = "forbidden_start_at"
	RoomRiskColForbiddenEndAt   = "forbidden_end_at"
)
const (
	RREColRoomID        = "room_id"
	RREColAuditDesc     = "audit_desc"
	RREColAuditInterval = "audit_interval"
	RREColLevel         = "level"
)

// 所有数据列
var RoomRiskColumns = []string{
	RoomRiskColRoomID,
	RoomRiskColEventTotal,
	RoomRiskColIgnoreTotal,
	RoomRiskColWarningTotal,
	RoomRiskColForbiddenTotal,
	RoomRiskColForbiddenStartAt,
	RoomRiskColForbiddenEndAt,
}

// 表名常量
const (
	TableRoomRisk      = "t_roomrisk"
	TableRoomRiskEvent = "t_roomriskevent"
)

// 表数据结构
type RoomInfo struct {
	RoomID           string    `json:"room_id"`            // 房间id
	EventTotal       uint32    `json:"event_total"`        // 总的事件次数
	IgnoreTotal      uint32    `json:"ignore_total"`       // 忽略事件次数
	WarningTotal     uint32    `json:"warning_total"`      // 警告次数
	ForbiddenTotal   uint32    `json:"forbidden_total"`    // 封禁次数
	ForbiddenStartAt time.Time `json:"forbidden_start_at"` // 封禁开始时间
	ForbiddenEndAt   time.Time `json:"forbidden_end_at"`   // 封禁结束时间
}

type RoomRiskEvent struct { // base room info
	ID            uint64    `json:"id"`             // ID
	RoomID        string    `json:"room_id"`        // 房间id
	Desc          string    `json:"desc"`           // 描述
	CreateAt      time.Time `json:"create_at"`      // 创建时间
	ToUserID      uint64    `json:"to_user_id"`     // 针对的用户（被举报的用户)
	FromUserID    uint64    `json:"from_user_id"`   // 风险来源用户
	AuditDesc     string    `json:"audit_desc"`     // 审核意见
	AuditInterval int64     `json:"audit_interval"` // 处理时长(s) , <0为永久
	RoomNum       string    `json:"room_num"`       // 房号
	RoomTitle     string    `json:"room_title"`     // 房间标题
	Level         int       `json:"level"`          // 风险级别  1 : 忽略(语音) 2 : 警告，低危 3 : 危险 ，高危 4 : 灾难(预留)
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
func (*RoomRiskEvent) TableName() string {
	return TableRoomRiskEvent
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
		case RoomRiskColForbiddenTotal:
			cols = append(cols, lf)
			params = append(params, e.ForbiddenTotal)
		case RoomRiskColForbiddenStartAt:
			cols = append(cols, lf)
			params = append(params, e.ForbiddenStartAt)
		case RoomRiskColForbiddenEndAt:
			cols = append(cols, lf)
			params = append(params, e.ForbiddenEndAt)
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
		case RoomRiskColEventTotal:
			if dbRoomInfo != nil && dbRoomInfo.EventTotal != e.ForbiddenTotal {
				cols = append(cols, lf)
				params = append(params, e.EventTotal)
				dbRoomInfo.EventTotal = e.EventTotal // 最后更新的值
			}
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
		case RoomRiskColForbiddenTotal:
			if dbRoomInfo != nil && dbRoomInfo.ForbiddenTotal != e.ForbiddenTotal {
				cols = append(cols, lf)
				params = append(params, e.ForbiddenTotal)
				dbRoomInfo.ForbiddenTotal = e.ForbiddenTotal
			}
		case RoomRiskColForbiddenStartAt:
			if dbRoomInfo != nil && dbRoomInfo.ForbiddenStartAt != e.ForbiddenStartAt {
				cols = append(cols, lf)
				params = append(params, e.ForbiddenStartAt)
				dbRoomInfo.ForbiddenStartAt = e.ForbiddenStartAt
			}
		case RoomRiskColForbiddenEndAt:
			if dbRoomInfo != nil && dbRoomInfo.ForbiddenEndAt != e.ForbiddenEndAt {
				cols = append(cols, lf)
				params = append(params, e.ForbiddenEndAt)
				dbRoomInfo.ForbiddenEndAt = e.ForbiddenEndAt
			}
		}
	}
	return cols, params
}

func (e *RoomRiskEvent) Updates(dbOldToNew Updater, fields []string) ([]string, []interface{}) {
	dbRRE, ok := dbOldToNew.(*RoomRiskEvent)
	if !ok {
		return nil, nil
	}
	cols := make([]string, 0)
	params := make([]interface{}, 0)
	for _, f := range fields {
		lf := strings.ToLower(f)
		switch lf {
		case RREColAuditDesc:
			if dbRRE != nil && dbRRE.AuditDesc != e.AuditDesc {
				cols = append(cols, lf)
				params = append(params, e.AuditDesc)
				dbRRE.AuditDesc = e.AuditDesc
			}
		case RREColRoomID:
			if dbRRE != nil && dbRRE.RoomID != e.RoomID {
				cols = append(cols, lf)
				params = append(params, e.RoomID)
				dbRRE.RoomID = e.RoomID
			}
		case RREColLevel:
			if dbRRE != nil && dbRRE.Level != e.Level {
				cols = append(cols, lf)
				params = append(params, e.Level)
				dbRRE.Level = e.Level
			}
		case RREColAuditInterval:
			if dbRRE != nil && dbRRE.AuditInterval != e.AuditInterval {
				cols = append(cols, lf)
				params = append(params, e.AuditInterval)
				dbRRE.AuditInterval = e.AuditInterval
			}

		}
	}
	return cols, params
}
