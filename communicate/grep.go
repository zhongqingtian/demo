package communicate

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	//create cmd
	cmd_go_env := exec.Command("go", "env")
	cmd_grep := exec.Command("grep", "GOROOT")

	stdout_env, env_error := cmd_go_env.StdoutPipe()
	if env_error != nil {
		fmt.Println("Error happened about standard output pipe ", env_error)
		return
	}

	//env_error := cmd_go_env.Start()
	if env_error := cmd_go_env.Start(); env_error != nil {
		fmt.Println("Error happened in execution ", env_error)
		return
	}
	/*
	   a1 := make([]byte, 1024)
	   n, err := stdout_env.Read(a1)
	   if err != nil {
	           fmt.Println("Error happened in reading from stdout", err)
	           return
	   }

	   fmt.Printf("Standard output of go env command: %s", a1[:n])
	*/
	//get the output of go env
	stdout_buf_grep := bufio.NewReader(stdout_env)

	//create input pipe for grep command
	stdin_grep, grep_error := cmd_grep.StdinPipe()
	if grep_error != nil {
		fmt.Println("Error happened about standard input pipe ", grep_error)
		return
	}

	//connect the two pipes together
	stdout_buf_grep.WriteTo(stdin_grep)

	//set buffer for reading
	var buf_result bytes.Buffer
	cmd_grep.Stdout = &buf_result

	//grep_error := cmd_grep.Start()
	if grep_error := cmd_grep.Start(); grep_error != nil {
		fmt.Println("Error happened in execution ", grep_error)
		return
	}

	err := stdin_grep.Close()
	if err != nil {
		fmt.Println("Error happened in closing pipe", err)
		return
	}

	//make sure all the infor in the buffer could be read
	if err := cmd_grep.Wait(); err != nil {
		fmt.Println("Error happened in Wait process")
		return
	}
	fmt.Println(buf_result.String())

}
