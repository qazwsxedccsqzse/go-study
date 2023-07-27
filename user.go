package main

import (
	"fmt"
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

	// new context
	context := Context{W: w, R: r}
	err := context.ReadJson(req)
	if err != nil {
		fmt.Fprintf(w, "context read json 失敗, %v", err)
		return
	}

	// 假設成功
	userMap := &map[string]int{"id": 1}
	context.WriteJson(http.StatusOK, userMap)
}
