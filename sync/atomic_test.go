package sync

import "testing"

func TestTAddInt(t *testing.T) {
	TAddInt()
}

func TestCAS(t *testing.T) {
	t.Log("old value = ", value)
	CAS(3)
	t.Log("CAS value = ", value)
}
