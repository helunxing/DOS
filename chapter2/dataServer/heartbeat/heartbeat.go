package heartbeat

import (
	"lib/rabbitmq"
	"os"
	"time"
)

// 建立消息队列结构体，发送监听地址
func StartHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}
}
