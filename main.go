package main

func main() {
	// 這個 main 函數就做兩件事情
	// 1. 註冊路由
	// 2. 監聽並且啟動 server

	// // http body 讀一次就消失
	// http.HandleFunc("/read_body_once", demo.ReadBodyOnce)
	// // http request.URL 只有 path 拿得到資料
	// http.HandleFunc("/read_url", demo.WholeUrl)
	// http.ListenAndServe(":3000", nil)

	// type demo
	// TestType()

	// struct demo
	// TestStruct()

	server := NewHttpServer("test server")
	// server.Route("/read_body_once", demo.ReadBodyOnce)
	// server.Route("/read_url", demo.WholeUrl)
	server.Route("/sign_up", SignUp)
	server.Start(":3000")
}
