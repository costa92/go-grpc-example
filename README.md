# go-grpc-example

## 生产 proto 文件

```
protoc --go_out=plugins=grpc:../server *.proto
```