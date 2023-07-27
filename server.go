package main

import "net/http"

type Server interface {
	Route(pattern string, handleFunc http.HandlerFunc)
	Start(address string) error
}

type SdkHttpServer struct {
}

// 註冊路由
func (s *SdkHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handleFunc)
}

// 啟動 server
func (s *SdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}
