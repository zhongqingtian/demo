package mp4

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	Run()
}

func TestVaint(t *testing.T) {
	buf := make([]byte, binary.MaxVarintLen64)

	for _, x := range []uint64{1, 2, 127, 128, 255, 256} {
		n := binary.PutUvarint(buf, x)
		fmt.Printf("%x\n", buf[:n])
	}
}
