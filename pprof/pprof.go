package main
import (
"os"
"runtime/pprof"
)
//创建写入 cup.out mem.out 文件
func main() {
	w, _ := os.Create("cpu.out")
	defer w.Close()
	pprof.StartCPUProfile(w)
	defer pprof.StopCPUProfile()

	w2, _ := os.Create("mem.out")
	defer w2.Close()
	defer pprof.WriteHeapProfile(w2)

	Sum(3, 5)

}

func Sum(a, b int) int {
	return a + b
}