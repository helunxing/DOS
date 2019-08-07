package versions

import (
	"encoding/json"
	"lib/es"
	"log"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	// 检查是否为get
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	from := 0
	size := 1000
	// 获取对象名
	name := strings.Split(r.URL.EscapedPath(), "/")[2]

	// 无限循环
	for {
		// 获取所有版本的元数据，获取size个
		metas, e := es.SearchAllVersions(name, from, size)
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// 元数据写入响应
		for i := range metas {
			b, _ := json.Marshal(metas[i])
			w.Write(b)
			w.Write([]byte("\n"))
		}
		// 已经读完则结束
		if len(metas) != size {
			return
		}
		// 未读完则继续获取
		from += size
	}
}
