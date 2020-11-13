package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

const html = `<html>
    <head>
    <meta charset="utf8">
    </head>
    <body>
        <form method="post" enctype="multipart/form-data">
            <input type="file" name="image" />
            <input type="submit" />
        </form>
    </body>
</html>`

func main() {
	http.HandleFunc("/upload/", uploadHandle)    // 上传
	http.HandleFunc("/uploaded/", showPicHandle) //显示图片
	err := http.ListenAndServe(":80", nil)
	fmt.Println(err)
}

// 上传图像接口
func uploadHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	req.ParseForm()
	if req.Method != "POST" {
		w.Write([]byte(html))
	} else {
		// 接收图片
		uploadFile, handle, err := req.FormFile("image")
		errorHandle(err, w)

		// 检查图片后缀
		ext := strings.ToLower(path.Ext(handle.Filename))
		if ext != ".jpg" && ext != ".png" {
			errorHandle(errors.New("只支持jpg/png图片上传"), w)
			return
			//defer os.Exit(2)
		}

		// 保存图片
		os.Mkdir("./uploaded/", 0777)
		saveFile, err := os.OpenFile("./uploaded/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666) //创建文件名

		errorHandle(err, w)
		io.Copy(saveFile, uploadFile) //把图片资源copy到刚刚创建的文件名下面
		defer uploadFile.Close()
		defer saveFile.Close()
		// 上传图片成功
		s := `<html>
    <head>
    <meta charset="utf8">
    </head>
    <body>
        查看上传图片: <a target='_blank' href="/uploaded/`
		s = s + handle.Filename
		s = s + `" >` + handle.Filename + `</a> </body>
</html>`
		w.Write([]byte(s))
	}
}

// 显示图片接口
func showPicHandle(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("." + req.URL.Path)
	errorHandle(err, w)
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	errorHandle(err, w)
	w.Write(buff)
}

// 统一错误输出接口
func errorHandle(err error, w http.ResponseWriter) {
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
