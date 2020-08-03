package project

import "testing"

func TestParseTableScanners(t *testing.T) {

	tableName := new(TableName)
	cols, scans := tableName.Scans("id", "start_time")
	t.Log(cols, scans[0])

}

func TestMakeWaitGroup(t *testing.T) {
	MakeWaitGroup()
}
