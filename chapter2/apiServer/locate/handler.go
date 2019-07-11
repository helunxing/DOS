package locate

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method

	// 不为get则返回不允许错误
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// 定位结果
	info := Locate(strings.Split(r.URL.EscapedPath(), "/")[2])
	// 失败则返回未找到
	if len(info) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// 返回节点地址
	b, _ := json.Marshal(info)
	w.Write(b)
}
