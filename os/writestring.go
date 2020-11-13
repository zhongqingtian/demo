package os

import (
	"fmt"
	"io"
	"os"
)

func OsFile2() {
	file, err := os.Create("./os/magic_msg.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	if _, err := io.WriteString(file, "Go is fun!"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
