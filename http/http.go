package main;

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.User.Username())
	w.Write([]byte("Hello"))
}
func say(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}
func main() {
	http.HandleFunc("/hello", hello);
	http.Handle("/handle", http.HandlerFunc(say));
	http.ListenAndServe(":8001", nil);
	fmt.Println(http.TimeFormat)
	select {}; // 阻塞进程
}
