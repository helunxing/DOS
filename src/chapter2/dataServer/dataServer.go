package main

import (
	"log"
	"net/http"
	"os"

	"../../chapter1/objects"
	"./heartbeat"
	"./locate"
	"./objects"
)

// 启动定位和心跳协程，绑定接口
func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
