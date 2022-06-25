# helloword

### proto 执行
```shell
protoc --go_out=services --go_opt=paths=source_relative \
    --go-grpc_out=services --go-grpc_opt=paths=source_relative \
    proto/hello.proto
```

但是生成的Server端接口中会出现一个mustEmbedUnimplemented***方法，好像是为了解决前向兼容问题，
不太懂，但是如果不解决，就无法传递给RegisterXXXService方法，解决办法有两个：

1. 在grpc server实现结构体中匿名嵌入Unimplemented***Server结构体

2. 使用protoc生成server代码时命令行加上关闭选项，protoc --go-grpc_out=require_unimplemented_servers=false

```shell
protoc --go_out=services --go_opt=paths=source_relative \
    --go-grpc_out=require_unimplemented_servers=false:services --go-grpc_opt=paths=source_relative \
    proto/hello.proto
```