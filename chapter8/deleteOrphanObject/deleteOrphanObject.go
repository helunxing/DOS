package main

import (
	"lib/es"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 获取目录下所有文件
	files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/*")
	// 查看元服务中是否有该hash值
	for i := range files {
		hash := strings.Split(filepath.Base(files[i]), ".")[0]
		hashInMetadata, e := es.HasHash(hash)
		if e != nil {
			log.Println(e)
			return
		}
		if !hashInMetadata {
			del(hash)
		}
	}
}

// 发出删除请求
func del(hash string) {
	log.Println("delete", hash)
	url := "http://" + os.Getenv("LISTEN_ADDRESS") + "/objects/" + hash
	request, _ := http.NewRequest("DELETE", url, nil)
	client := http.Client{}
	client.Do(request)
}
