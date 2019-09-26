package main

import (
	"lib/es"
	"log"
)

const MIN_VERSION_COUNT = 5

func main() {
	// 搜索版本数量大于6的对象
	buckets, e := es.SearchVersionStatus(MIN_VERSION_COUNT + 1)
	if e != nil {
		log.Println(e)
		return
	}
	for i := range buckets {
		bucket := buckets[i]
		// 删除至只剩五个
		for v := 0; v < bucket.Doc_count-MIN_VERSION_COUNT; v++ {
			es.DelMetadata(bucket.Key, v+int(bucket.Min_version.Value))
		}
	}
}
