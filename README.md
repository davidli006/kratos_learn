# kratos学习

## protoc安装
- https://github.com/protocolbuffers/protobuf/releases 下载地址

## protoc-gen-go安装
- go get github.com/golang/protobuf/protoc-gen-go

## kratos安装
- go install github.com/go-kratos/kratos/cmd/kratos/v2@latest

## 创建项目
- cd ./backend
- kratos new verifyCode
- go mod tidy
- go get github.com/google/wire/cmd/wire
- go generate ./...
- 


