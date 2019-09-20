package temp

import (
	"os"

	"../locate"
)

// 将文件移动到对象文件夹以转正
func commitTempObject(datFile string, tempinfo *tempInfo) {
	os.Rename(datFile, os.Getenv("STORAGE_ROOT")+"/objects/"+tempinfo.Name)
	locate.Add(tempinfo.Name)
}
