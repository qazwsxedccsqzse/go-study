package main

import "net/http"

type Server interface {
	// method GET, POST, PUT, DELETE, ..
	Route(method string, pattern string, handleFunc func(c *Context))
	Start(address string) error
}

type SdkHttpServer struct {
	name string
	// 註冊的路由
	handler *HandlerBasedOnMap
}

// 註冊路由
func (s *SdkHttpServer) Route(method string, pattern string, handleFunc func(c *Context)) {
	key := s.handler.Key(method, pattern)
	s.handler.handlers[key] = handleFunc
	// http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
	// 	context := NewContext(w, r)
	// 	handleFunc(context)
	// })
}

// 啟動 server
func (s *SdkHttpServer) Start(address string) error {
	// 因為最後只註冊一次, 只需要在 Server 啟動的時候才註冊

	http.Handle("/", s.handler)

	return http.ListenAndServe(address, nil)
}

// 要返回 interface instance 的時候, 內部返回必為指標
func NewHttpServer(name string) Server {
	server := &SdkHttpServer{name: name}

	return server
}
