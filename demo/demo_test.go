package demo

import (
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
	FixBug()
}

func TestGoRun(t *testing.T) {
	GoRun()
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