package locate

import (
	"encoding/json"
	"lib/rabbitmq"
	"lib/rs"
	"lib/types"
	"os"
	"time"
)

func Locate(name string) (locateInfo map[int]string) {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	q.Publish("dataServers", name)
	// 超时则关闭channel
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	locateInfo = make(map[int]string)
	// 最多收取M+N条消息
	for i := 0; i < rs.ALL_SHARDS; i++ {
		msg := <-c
		if len(msg.Body) == 0 {
			return
		}
		var info types.LocateMessage
		json.Unmarshal(msg.Body, &info)
		locateInfo[info.Id] = info.Addr
	}
	return
}

// 判断消息是否满足读取条件
func Exist(name string) bool {
	return len(Locate(name)) >= rs.DATA_SHARDS
}
