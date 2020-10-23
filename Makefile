export GOPROXY=https://goproxy.cn

build: cmd/main.go
	mkdir -p build/bin
	go build cmd/main.go && mv main build/bin

codegen: proto/echo.proto
	mkdir -p proto/gen
	protoc -I./proto -I. --go_out=plugins=grpc,paths=source_relative:proto/gen/ $<
	protoc -I./proto -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:proto/gen/ $<
	protoc -I./proto -I. --swagger_out=logtostderr=true:proto/gen $<
