package objects

import (
	"io"
	"os"
)

// 打开文件并写入
func sendFile(w io.Writer, file string) {
	f, _ := os.Open(file)
	defer f.Close()
	io.Copy(w, f)
}
