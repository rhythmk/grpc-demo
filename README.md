# grpc-demo


## 安装

-  **初始项目** 
 go mod init grpc-demo

- **安装 grpc包**
```
 go get google.golang.org/grpc
```
- **定义文件**

定义 proto 文件，具体可以参考：

>Protobuf 终极教程
https://colobu.com/2019/10/03/protobuf-ultimate-tutorial-in-go/

- **生成对应代码**

```
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

通过命令，生成pb ，以及grpc对应的代码：

```
protoc --go_out=.  proto/helloservice.proto
protoc --go-grpc_out=.   proto/helloservice.proto 

```

## 运行代码

分别运行 :
```
go run  .\server\main.go
go run  .\client\main.go

```


##  常见异常

- **异常1**

 Please specify either:
        • a "go_package" option in the .proto source file, or
        • a "M" argument on the command line.

解决方案:
> proto 中指定 ` option go_package = "./hellogrpc;hellogrpc"; `
go_package 由两部分组成，中间用逗号分隔。 前面为生成文件路径，后面为生成文件包名。


- **异常2**

实现service 过程会出现。
 missing method mustEmbedUnimplementedHelloServiceServer

 解决方案 :
 >继承 xxx.UnimplementedHelloServiceServer  结构体 ：
代码如下：
```go
 
 type RpcHelloServiceServer struct {
	*hellogrpc.UnimplementedHelloServiceServer
 }

 ```