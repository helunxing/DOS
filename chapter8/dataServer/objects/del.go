package objects

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"../locate"
)

// 根据hash移除文件
func del(w http.ResponseWriter, r *http.Request) {
	hash := strings.Split(r.URL.EscapedPath(), "/")[2]
	files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/" + hash + ".*")
	if len(files) != 1 {
		return
	}
	// 移出缓存
	locate.Del(hash)
	// 移到待删除目录，但真删除前，需检查是否存在
	os.Rename(files[0], os.Getenv("STORAGE_ROOT")+"/garbage/"+filepath.Base(files[0]))
}
