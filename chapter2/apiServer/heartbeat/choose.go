package heartbeat

import (
	"math/rand"
)

//随机选择发送目的节点
func ChooseRandomDataServer() string {
	ds := GetDataServers()
	n := len(ds)
	if n == 0 {
		return ""
	}
	return ds[rand.Intn(n)]
}
