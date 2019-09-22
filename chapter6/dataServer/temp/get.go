package temp

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// 打开并将uuid作为响应正文输出
func get(w http.ResponseWriter, r *http.Request) {
	uuid := strings.Split(r.URL.EscapedPath(), "/")[2]
	f, e := os.Open(os.Getenv("STORAGE_ROOT") + "/temp/" + uuid + ".dat")
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
