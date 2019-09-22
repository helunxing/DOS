package objects

import (
	"fmt"
	"io"
	"lib/es"
	"lib/utils"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//
func get(w http.ResponseWriter, r *http.Request) {
	// 文件名和版本
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	versionId := r.URL.Query()["version"]
	version := 0
	var e error
	if len(versionId) != 0 {
		version, e = strconv.Atoi(versionId[0])
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	// 获取元数据
	meta, e := es.GetMetadata(name, version)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if meta.Hash == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	hash := url.PathEscape(meta.Hash)
	// 获取复原后的输出流
	stream, e := GetStream(hash, meta.Size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// 从请求头获取偏移量
	offset := utils.GetOffsetFromHeader(r.Header)
	if offset != 0 {
		// 将流跳转到该位置
		stream.Seek(offset, io.SeekCurrent)
		// 写入范围头
		w.Header().Set("content-range", fmt.Sprintf("bytes %d-%d/%d", offset, meta.Size-1, meta.Size))
		w.WriteHeader(http.StatusPartialContent)
	}
	// 复制
	io.Copy(w, stream)
	stream.Close()
}
