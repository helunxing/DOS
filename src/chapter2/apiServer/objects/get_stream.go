package objects

import (
	"fmt"
	"io"
	"lib/objectstream"

	"../locate"
)

// 发送查询请求，若失败则返回错误并输出
func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return objectstream.NewGetStream(server, object)
}
