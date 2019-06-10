package main

import (
	"log"
	"net/http"
	"os"

	"./objects"
)

//绑定接口并运行
func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}

// 简单地实现了写入和读取
// 还没有体现分布式，也没有体现对象存储
