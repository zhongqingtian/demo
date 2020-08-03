package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {
	proverbs := []string {
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}

	for _, p := range proverbs{
		n, err := os.Stdout.Write([]byte(p)) //打印到控制台
		time.Sleep(1*time.Second)//隔一秒发一次
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if n != len(p) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}


}