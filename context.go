package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ReadJson(req interface{}) error {
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		return nil
	}

	// 把讀進來的 json 對應到 sign up request
	err = json.Unmarshal(body, req)
	if err != nil {
		return nil
	}

	return nil
}

func (c *Context) WriteJson(code int, data interface{}) error {
	response := &CommonResponse{
		Success: true,
		Msg:     "註冊成功",
		Data:    data,
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		return err
	}

	// 設定 content-type header
	c.W.Header().Set("Content-Type", "application/json")
	// 寫入狀態碼
	c.W.WriteHeader(code)

	_, err = c.W.Write(responseJson)

	return err
}

func (c *Context) SuccessJson(data interface{}) error {
	return c.WriteJson(http.StatusOK, data)
}

func (c *Context) BadRequestJson(data interface{}) error {
	return c.WriteJson(http.StatusBadRequest, data)
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{W: w, R: r}
}
