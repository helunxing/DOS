package objects

import (
	"io"
	"log"
	"net/http"
	"strings"
)

// 将获取到的对象则写入http相应，否则返回未找到
func get(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	stream, e := getStream(object)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	io.Copy(w, stream)
}
