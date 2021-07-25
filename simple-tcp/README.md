# simple-tcp

### 目录结构：
```shell
$ tree simple-tcp
├── README.md   
├── client
│   └── client.go
├── proto
│   └── search.proto
├── server
│   └── search.pb.go
└── server.go

```

在 proto 执行下面的命令 

```shell
protoc --go_out=plugins=grpc:../server *.proto
```
plugins=plugin1+plugin2：指定要加载的子插件列表

我们定义的 proto 文件是涉及了 RPC 服务的，而默认是不会生成 RPC 代码的，因此需要给出 plugins 参数传递给 protoc-gen-go，告诉它，请支持 RPC（这里指定了 gRPC）

–go_out=.：设置 Go 代码输出的目录

../server 是指生成的文件放在 server 目录下

该指令会加载 protoc-gen-go 插件达到生成 Go 代码的目的，生成的文件以 .pb.go 为文件后缀

: （冒号）
冒号充当分隔符的作用，后跟所需要的参数集。如果这处不涉及 RPC，命令可简化为：
```shell
$ protoc --go_out=. *.proto
```

执行 服务端
```shell
go run server.go
```

执行客户端

```shell
go run client/client.go
```
执行结果：

```shell
$ go run client/client.go
get test Server
```