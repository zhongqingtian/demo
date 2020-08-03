package main

import (
	"fmt"
	"strings"
)

func shsh()  {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
}

func main()  {
	shsh()
}