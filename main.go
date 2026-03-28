package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter,r *http.Request){
	fmt.Println("收到了一次访问")
	w.Write([]byte("Hello,this is ping route"))
}

func main(){
	http.HandleFunc("/ping",pingHandler)
	fmt.Println("服务器正在启动，监听端口 8080...")
	http.ListenAndServe(":8080", nil)
}