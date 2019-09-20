package objects

import (
	"lib/es"
	"lib/utils"
	"log"
	"net/http"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	// 获得hash值
	hash := utils.GetHashFromHeader(r.Header)
	if hash == "" {
		log.Println("missing object hash in digest header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 写入
	size := utils.GetSizeFromHeader(r.Header)
	c, e := storeObject(r.Body, hash, size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(c)
		return
	}
	// 返回错误码
	if c != http.StatusOK {
		w.WriteHeader(c)
		return
	}

	// 将名称写入es
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	e = es.AddVersion(name, hash, size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
