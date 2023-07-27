package main

import (
	"net/http"
)

type HandlerBasedOnMap struct {
	// 已註冊的 handler
	// key = http method + "#" + url
	// value = 註冊對應的 handler func
	handlers map[string]func(ctx *Context)
}

// 要 implement net/http/server.go 內的 ServeHTTP interface

func (h *HandlerBasedOnMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := h.Key(r.Method, r.URL.Path)

	// 先檢查該路由有沒有註冊過, 如果是 ok 就代表已經註冊過了
	if handlerFunc, ok := h.handlers[key]; ok {
		handlerFunc(NewContext(w, r))
	} else {
		// 沒找到
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	}
}

// 取得這個註冊路由對應的 key
func (h *HandlerBasedOnMap) Key(method, pattern string) string {
	return method + "#" + pattern
}
