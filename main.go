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

	var num1, num2,realResult int
	var operator string
	_, err = fmt.Sscanf(req.Expression, "%d %s %d", &num1, &operator, &num2)
	
	if err != nil {
		fmt.Println("解析表达式失败:", err)
		return 
	} 

	switch operator{
	case "+":realResult=num1+num2
	case "-":realResult=num1-num2
	case "*":realResult=num1*num2
	case "/":realResult=num1/num2
	default:
		http.Error(w, "不支持的运算符", http.StatusBadRequest)
		return
	}
	resp:=CalcResponse{
		Answer:realResult,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main(){
	http.HandleFunc("/ping",pingHandler)
	fmt.Println("服务器正在启动，监听端口 8080...")
	http.ListenAndServe(":8080", nil)
}