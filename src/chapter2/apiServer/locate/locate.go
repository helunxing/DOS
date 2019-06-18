package locate

import (
	"lib/rabbitmq"
	"os"
	"strconv"
	"time"
)

// 向数据节点发送查询请求，并接受
func Locate(name string) string {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))

	//发送查询请求。
	// 问题：如何实现建立临时队列，并接收数据节点的返回包
	q.Publish("dataServers", name)
	c := q.Consume()

	// 超时即关闭
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	// 若关闭默认将返回空字符串
	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

// 判断要定位的文件是否存在
func Exist(name string) bool {
	return Locate(name) != ""
}
