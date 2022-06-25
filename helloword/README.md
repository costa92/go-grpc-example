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

```golang
type HelloServiceImpl struct {
    helloword.UnimplementedHelloServiceServer
}
```
2. 使用protoc生成server代码时命令行加上关闭选项，protoc --go-grpc_out=require_unimplemented_servers=false

```shell
protoc --go_out=services --go_opt=paths=source_relative \
    --go-grpc_out=require_unimplemented_servers=false:services --go-grpc_opt=paths=source_relative \
    proto/hello.proto
```

### 参考: https://blog.csdn.net/guolianggsta/article/details/118075613?spm=1001.2101.3001.6650.13&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-13-118075613-blog-108135724.pc_relevant_downloadblacklistv1&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-13-118075613-blog-108135724.pc_relevant_downloadblacklistv1&utm_relevant_index=16