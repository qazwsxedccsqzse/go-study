package demo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 用來 demo r.Body 只能讀取一次
func ReadBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		return
	}

	fmt.Fprintf(w, "第一次讀取: %s", string(body))

	// 再次讀取
	body, err = io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "再次讀取失敗: %v", err)
		return
	}
	fmt.Fprintf(w, "第二次讀取: %s", string(body))
}

// r.URL 只有 path 不會為空
func WholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
}
