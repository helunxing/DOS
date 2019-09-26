#!/bin/bash
# 上传测试文件
curl -v 10.29.2.1:12345/objects/test7 -XPUT --data-binary @/tmp/file -H "Digest: SHA-256=IEkqTQ2E+L6xdn9mFiKfhdRMKCe2S9v7Jg7hL6EQng4="
# 下载文件
curl -v 10.29.2.1:12345/objects/test7 -o /tmp/output
# 比较
diff -s /tmp/output /tmp/file
# 查看分片
ls -ltr /tmp/?/objects
# 接受压缩
curl -v 10.29.2.1:12345/objects/test7 -H "Accept-Encoding: gzip" -o /tmp/output2.gz

gunzip /tmp/output2.gz
# 比较
diff -s /tmp/output2 /tmp/file
