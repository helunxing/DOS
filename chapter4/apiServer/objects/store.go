package objects

import (
	"fmt"
	"io"
	"lib/utils"
	"net/http"
	"net/url"

	"../locate"
)

// 需要确定对象大小，多了size参数
func storeObject(r io.Reader, hash string, size int64) (int, error) {
	if locate.Exist(url.PathEscape(hash)) {
		return http.StatusOK, nil
	}
	// 确保hash值安全
	stream, e := putStream(url.PathEscape(hash), size)
	if e != nil {
		return http.StatusInternalServerError, e
	}
	// reader读取时，读取自r，同时写入stream
	reader := io.TeeReader(r, stream)

	d := utils.CalculateHash(reader)
	if d != hash {
		stream.Commit(false)
		return http.StatusBadRequest, fmt.Errorf("object hash mismatch, calculated=%s, requested=%s", d, hash)
	}
	stream.Commit(true)
	return http.StatusOK, nil
}
