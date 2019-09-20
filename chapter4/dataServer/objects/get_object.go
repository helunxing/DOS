package objects

import (
	"lib/utils"
	"log"
	"net/url"
	"os"

	"../locate"
)

// 检查文件hash并返回其路径，防止数据存储过程中出错
func getFile(hash string) string {
	file := os.Getenv("STORAGE_ROOT") + "/objects/" + hash
	f, _ := os.Open(file)
	d := url.PathEscape(utils.CalculateHash(f))
	f.Close()
	if d != hash {
		log.Println("object hash mismatch, remove", file)
		locate.Del(hash)
		os.Remove(file)
		return ""
	}
	return file
}
