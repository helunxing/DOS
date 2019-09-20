package objects

import (
	"net/http"
	"strings"
)

// 取hash值，取文件，返回文件
func get(w http.ResponseWriter, r *http.Request) {
	file := getFile(strings.Split(r.URL.EscapedPath(), "/")[2])
	if file == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// 写入响应
	sendFile(w, file)
}
