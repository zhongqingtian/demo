package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed ss
var s string
//go:embed demo
var b []byte

//go:embed template
var fs embed.FS
// https://www.flysnow.org/2021/02/28/golang-embed-for-web.html
func main() {
	fmt.Println(s)
	// err := os.WriteFile("D:\\goCode\\src\\demo\\instruct\\demo",[]byte(s),os.ModePerm)
	//fmt.Println(err)
	fmt.Println(string(b))
	data, err := fs.ReadFile("template/index.html")
	fmt.Println(err,string(data))

}
