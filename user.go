package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 定義請求
type SignUpRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

// 定義 response
type CommonResponse struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	// 先產生 sign up request
	req := new(SignUpRequest)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "讀取 body 失敗, %v", err)
		return
	}

	// 把讀進來的 json 對應到 sign up request
	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Fprintf(w, "json 轉換失敗, %v", err)
		return
	}

	// 假設成功
	userMap := map[string]int{"id": 1}
	response := CommonResponse{
		Success: true,
		Msg:     "註冊成功",
		Data:    userMap,
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		fmt.Fprintf(w, "json encode失敗, %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(responseJson))
}
