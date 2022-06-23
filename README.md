# go-grpc-example

## 生产 proto 文件


生成代码

```shell
protoc --go_out=../services Prod.proto  
```
生成grpc 服务

```shell
protoc --go_out=plugins=grpc:../server *.proto
```