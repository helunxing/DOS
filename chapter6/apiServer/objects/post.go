package objects

import (
	"lib/es"
	"lib/rs"
	"lib/utils"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"../heartbeat"
	"../locate"
)

// 对对象请求返回token
func post(w http.ResponseWriter, r *http.Request) {
	// 获取对象名，大小
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	size, e := strconv.ParseInt(r.Header.Get("size"), 0, 64)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// 获取hash
	hash := utils.GetHashFromHeader(r.Header)
	if hash == "" {
		log.Println("missing object hash in digest header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 存在则添加新版本
	if locate.Exist(url.PathEscape(hash)) {
		e = es.AddVersion(name, hash, size)
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		return
	}
	// 不存在则选写点写入
	ds := heartbeat.ChooseRandomDataServers(rs.ALL_SHARDS, nil)
	if len(ds) != rs.ALL_SHARDS {
		log.Println("cannot find enough dataServer")
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	// 返回RSesustream，其中含有token(uuid,名字大小hash节点地址)和rsputstream
	stream, e := rs.NewRSResumablePutStream(ds, name, url.PathEscape(hash), size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 写入base编码后的json格式内容
	w.Header().Set("location", "/temp/"+url.PathEscape(stream.ToToken()))
	w.WriteHeader(http.StatusCreated)
}
