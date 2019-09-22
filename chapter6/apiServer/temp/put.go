package temp

import (
	"io"
	"lib/es"
	"lib/rs"
	"lib/utils"
	"log"
	"net/http"
	"net/url"
	"strings"

	"../locate"
)

func put(w http.ResponseWriter, r *http.Request) {
	// 根据token获得文件流
	token := strings.Split(r.URL.EscapedPath(), "/")[2]
	stream, e := rs.NewRSResumablePutStreamFromToken(token)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// 获取token当前大小
	current := stream.CurrentSize()
	if current == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// 需要与偏移量相同
	offset := utils.GetOffsetFromHeader(r.Header)
	if current != offset {
		w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
		return
	}
	bytes := make([]byte, rs.BLOCK_SIZE)
	for {
		// 以缓冲区长度为循环读取正文写入stream
		n, e := io.ReadFull(r.Body, bytes)
		if e != nil && e != io.EOF && e != io.ErrUnexpectedEOF {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// 累加上传量
		current += int64(n)
		// 超过对象大小，说明上传数据有误
		if current > stream.Size {
			stream.Commit(false)
			log.Println("resumable put exceed size")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		// 读到的长度不到缓冲区长，且总长度不等于对象长，说明本次已经结束
		// 丢弃，下次重传
		if n != rs.BLOCK_SIZE && current != stream.Size {
			return
		}
		// 向流中写入
		stream.Write(bytes[:n])
		// 上传了全部数据后，写入对象，并计算hash值。
		if current == stream.Size {
			stream.Flush()
			getStream, e := rs.NewRSResumableGetStream(stream.Servers, stream.Uuids, stream.Size)
			hash := url.PathEscape(utils.CalculateHash(getStream))
			if hash != stream.Hash {
				stream.Commit(false)
				log.Println("resumable put done but hash mismatch")
				w.WriteHeader(http.StatusForbidden)
				return
			}
			// hash存在则删除对象
			if locate.Exist(url.PathEscape(hash)) {
				stream.Commit(false)
			} else {
				stream.Commit(true)
			}
			// 添加版本
			e = es.AddVersion(stream.Name, stream.Hash, stream.Size)
			if e != nil {
				log.Println(e)
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	}
}
