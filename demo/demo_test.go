package demo

import (
	"fmt"
	"testing"
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

func TestFixBug(t *testing.T) {
	/*ll := []string{"kkk", "1111"}
	var s string*/
	 data := struct {
		 Name string
		 Age int64
	 }{}
	 data.Age = 10
	 data.Name = "11"
	 da := data
	 fmt.Println(da)
	/*go func() {
		for {
			n := rand.Int()
			s = ll[n%2]
		}
	}()

	for {
		go func(ss string) {
			/*defer func() {
				if err := recover(); err != nil {
					buf := make([]byte, 1024)
					buf = buf[:runtime.Stack(buf, true)]
					t.Log(buf)
				}
			}()*/
			// fmt.Println(fmt.Sprintf("%P",s))
		/*	_ = []byte(ss)
		}(s)
	}*/

}

func TestGoRun(t *testing.T) {
	GoRun()
}
