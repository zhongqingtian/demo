package demo

import "testing"

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
