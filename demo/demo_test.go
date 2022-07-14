package demo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestGetPerm(t *testing.T) {
	t.Log(GetPerm())
}

func TestForFiles(t *testing.T) {
	ForFiles("F:\\goCode\\src\\demo")
}

func TestForSelect(t *testing.T) {
	ch := make(chan int8, 1)
	ch <- 1
	t.Log(ForSelect(ch))
	for i := range ch {
		t.Log(i)
	}
	defer close(ch)
}

func TestName(t *testing.T) {
	// 	FixBug()
	/*str := "kk21"
	t.Log(strings.Split(str, ","))*/

	 m := make(map[int]int)
	// m[0]=0
	t.Log(len(m))
}

func TestEGroup(t *testing.T) {
	EGroup()
}

func TestSyncWait(t *testing.T) {
	SyncWait()
}

func TestGoRun(t *testing.T) {
	GoRun()
}


func TestSyncMap(t *testing.T) {
	SyncMap()
}
func TestFixBug(t *testing.T) {
	str := "anbc中国" // 中文占三个字节 len是字节长度
	fmt.Println(strings.Count(str,"")-1) // 6
	fmt.Println(len(str))// 10
	fmt.Println(utf8.RuneCountInString(str))// 6
}

func TestSS_GetName(t *testing.T) {
	a := SS{Name:"qq"}
	f1 := SS.GetName
	fmt.Println(f1(a))

     f2 := a.GetName()
     fmt.Println(f2)

     /// interface
     var e interface{}
     fmt.Println(reflect.TypeOf(e))
     f,_ := os.Open("egg.txt")
     e =f
	fmt.Println(reflect.TypeOf(e))
  var rw io.ReadWriter
     rw = f
	fmt.Println(reflect.TypeOf(rw))

}

type Much struct {
	Status         int    `json:"status"`
	Result         int    `json:"result"`
	ExpirationTime string `json:"expiration_time"`
	UserRole       int    `json:"user_role"`
	Desc           string `json:"desc"`
}

func Example_much() {
	demo := Much{}
	outInterface(context.Background(), demo, setMuch(&demo))

	// Output:
	// pointer_test.Much{Status:0, Result:0, ExpirationTime:"", UserRole:0, Desc:""}
}

func setMuch(demo *Much) error {
	return json.Unmarshal([]byte(`{"status":10,"result":1,"UserRole":6}`), demo)
}

type Single struct {
	Id int `json:"id"`
}

func Example_single() {
	demo := Single{}
	outInterface(context.Background(), demo, setSingle(&demo))

	// Output:
	// pointer_test.Single{Id:10}
}

func setSingle(demo *Single) error {
	return json.Unmarshal([]byte(`{"id":10}`), demo)
}

func outInterface(ctx context.Context, demo interface{}, err error)  {
	fmt.Printf("%#v", demo)
}