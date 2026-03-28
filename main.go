package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type CalcRequest struct {
	Expression string `json:"expression"` 
}

type CalcResponse struct {
	Answer int `json:"answer"`
}

func pingHandler(w http.ResponseWriter,r *http.Request){
	var req CalcRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "请求格式不对", http.StatusBadRequest)
		return
	}
	fmt.Println("成功获取前端请求：", req.Expression)
	fakeResult := 10000
	resp := CalcResponse{
		Answer: fakeResult,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main(){
	http.HandleFunc("/ping",pingHandler)
	fmt.Println("服务器正在启动，监听端口 8080...")
	http.ListenAndServe(":8080", nil)
}