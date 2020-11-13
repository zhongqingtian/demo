package options

import "testing"

func TestCollectOptions(t *testing.T) {
	CollectOptions(GetOptions1(), GetOptions2(), GetOptions3())
}
