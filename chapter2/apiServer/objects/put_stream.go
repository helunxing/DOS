package objects

import (
	"fmt"
	"lib/objectstream"

	"../heartbeat"
)

// 随机选择服务器并将发送文件
func putStream(object string) (*objectstream.PutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return objectstream.NewPutStream(server, object), nil
}
