package json

import (
	"testing"
)

func TestJsonStruct_Load(t *testing.T) {
	JsonParse := NewJsonStruct()
	v := Config{}
	//下面使用的相对路径，config.json文件
	JsonParse.Load("./config.json",&v)
	t.Log(v.Addr)
	t.Log(v.MogoCol.MongoDb)
}
