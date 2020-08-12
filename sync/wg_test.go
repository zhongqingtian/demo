package sync

import "testing"

func TestWgGro(t *testing.T) {
	t.Log(WgGro())
	it := make([]interface{}, 0)
	it = append(it, 1)
	st := []string{"a", "b"}
	it = append(it, "djdj", "dj")
	it = append(it, st)

}
