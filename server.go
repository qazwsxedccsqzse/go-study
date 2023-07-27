package main

import "net/http"

type Server interface {
	Route(pattern string, handleFunc func(c *Context))
	Start(address string) error
}

type SdkHttpServer struct {
	name string
}

// 註冊路由
func (s *SdkHttpServer) Route(pattern string, handleFunc func(c *Context)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		context := Context{W: w, R: r}
		handleFunc(&context)
	})
}

// 啟動 server
func (s *SdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

// 要返回 interface instance 的時候, 內部返回必為指標
func NewHttpServer(name string) Server {
	return &SdkHttpServer{name: name}
}
