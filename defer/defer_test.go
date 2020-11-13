package _defer

import (
	"testing"
	"time"
)

func TestGetReturn(t *testing.T) {
	// GetReturn()

	no1 := time.Now().UnixNano() / 1e6
	no2 := time.Now().Add(3*time.Minute).UnixNano() / 1e6
	t.Log(no1)
	t.Log(no2)
}
