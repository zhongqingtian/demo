package demo

import (
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