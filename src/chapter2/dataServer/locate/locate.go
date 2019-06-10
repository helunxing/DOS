package locate

import (
	"lib/rabbitmq"
	"os"
	"strconv"
)

// 访问文件名，判断其是否存在
func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("dataServers")

	// 消费者channel
	c := q.Consume()
	for msg := range c {
		// 去除JSON编码的双引号
		object, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		// 本节点存在则返回监听地址
		if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) {
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
	}
}
