package main

import (
	"lib/es"
	"lib/utils"
	"log"
	"os"
	"path/filepath"
	"strings"

	"../apiServer/objects"
)

func main() {
	// 获取所有文件
	files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/*")
	// 逐个检查数据
	for i := range files {
		hash := strings.Split(filepath.Base(files[i]), ".")[0]
		verify(hash)
	}
}

// 根据hash检查修复项目
func verify(hash string) {
	log.Println("verify", hash)
	// 搜索其大小
	size, e := es.SearchHashSize(hash)
	if e != nil {
		log.Println(e)
		return
	}
	// 创建对象数据流
	stream, e := objects.GetStream(hash, size)
	if e != nil {
		log.Println(e)
		return
	}
	// 计算hash值
	d := utils.CalculateHash(stream)
	if d != hash {
		log.Printf("object hash mismatch, calculated=%s, requested=%s", d, hash)
	}
	stream.Close()
}
