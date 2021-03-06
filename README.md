# grpc-gateway-demo
##### GO开发环境

```shell script
GOROOT=/usr/local/Cellar/go/1.15.3/libexec
export GOROOT
export PATH=$PATH:$GOROOT/bin
export GOPATH=$HOME/gopath
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```

##### 安装依赖库

```shell script
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```
#### 创建项目

```shell script
// github上创建空项目
git clone https://github.com/realwrtoff/grpc-gateway-demo.git
cd grpc-gateway-demo
// mod 管理
git mod init github.com/realwrtoff/grpc-gateway-demo
// grpc-gateway环境, 注意替换版本号
mkdir -p ./proto/google/api
cp $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.15.2/third_party/googleapis/google/api/* ./proto/google/api
cd proto
vim echo.proto
```
#### 编辑echo.proto文件

```protobuf
syntax = "proto3";
package echo;

option go_package = "github.com/realwrtoff/grpc-gateway-demo/proto/echo";

import "google/api/annotations.proto";

message EchoReq {
  string name = 1;
  int32 age = 2;
}

message EchoRes {
  string name = 1;
  int32 age = 2;
}

service EchoService {
  rpc Echo(EchoReq) returns (EchoRes) {
    option (google.api.http) = {
      get: "/v1/example/echo"
    };
  }
}

message Info {
  string op = 1;
  int64 a = 2;
  int64 b = 3;
}

message CalReq {
  string uid = 1;
  Info info = 2;
}

message CalRes {
  string uid = 1;
  int64 result = 2;
}

service CalService {
  rpc Cal(CalReq) returns (CalRes) {
    option (google.api.http) = {
      post: "/v2/example/cal/{uid}",
      body: "info"
    };
  }
}
```
#### 生成代码
```shell script
mkdir -p ./gen
protoc --go_out=plugins=grpc,paths=source_relative:./gen echo.proto
protoc --grpc-gateway_out=logtostderr=true,paths=source_relative:./gen echo.proto
protoc --swagger_out=logtostderr=true:./gen echo.proto
```
### 编写代码，编译
```shell script
// 代码参考 cmd/main.go
go mod tidy
go mod vendor
go build cmd/main.go
```
#### 运行测试
```shell script
./main &
curl "http://127.0.0.1:21680/v1/example/echo?name=jim&age=23"
// expect {"name":"jim","age":24}
curl -XPOST -d '{"a":1,"b":2,"op":"+"}' "http://127.0.0.1:21680/v2/example/cal/jim"
{"uid":"jim","result":"3"}
// expect {"uid":"jim","result":"3"}
```
