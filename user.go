package main

import (
	"fmt"
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

func SignUp(context *Context) {
	// 先產生 sign up request
	req := new(SignUpRequest)

	err := context.ReadJson(req)
	if err != nil {
		context.BadRequestJson(err)
		return
	}

	// 假設成功
	userMap := &map[string]int{"id": 1}
	err = context.SuccessJson(userMap)
	if err != nil {
		fmt.Printf("context write json 失敗, %s", err.Error())
		return
	}
}
