package heartbeat

import (
	"lib/rabbitmq"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

//收到心跳时间
var dataServers = make(map[string]time.Time)

//心跳写入锁
var mutex sync.Mutex

func ListenHeartbeat() {
	//消费数据服务的心跳信号
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("apiServers")
	c := q.Consume()

	//启动协程，用于移除掉线数据节点
	go removeExpiredDataServer()
	//更新收到数据服务心跳信息
	for msg := range c {
		dataServer, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

//遍历数据节点，十秒以上没有服务的即移除
func removeExpiredDataServer() {
	for {
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for s, t := range dataServers {
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}
		mutex.Unlock()
	}
}

//生成数据节点列表
func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string, 0)
	for s, _ := range dataServers {
		ds = append(ds, s)
	}
	return ds
}

//随机选择发送目的节点
func ChooseRandomDataServer() string {
	ds := GetDataServers()
	n := len(ds)
	if n == 0 {
		return ""
	}
	return ds[rand.Intn(n)]
}
