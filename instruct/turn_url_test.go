package instruct

import (
	"fmt"
	"testing"
)

func TestTurnUrls(t *testing.T) {
	m := TurnUrls(MapStruct{
		Name: "zsf",
		Age:  100,
		F64:  121.00000000000021,
	})
	t.Log(m.Encode())
}

func BenchmarkTurnUrls(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		TurnUrls(MapStruct{
			Name: "zsf",
			Age:  100,
			F64:  121.00000000000021,
		})
	}
}

func TestGooleTurnUrl(t *testing.T) {
	o :=MapStruct{
		Name: "zsf",
		Age:  100,
		F64:  121.00000000000021,
		//Ll: []int64{12,15,14},
		//Times: time.Now(),
	}
	GooleTurnUrls(o)
}

func TestConvertToMap(t *testing.T) {
	o :=MapStruct{
		Name: "zsf",
		Age:  100,
		F64:  121.00000000000021,
		//Ll: []int64{12,15,14},
		//Times: time.Now(),
	}
	m := ConvertToMap(o)
	fmt.Println(m)
}