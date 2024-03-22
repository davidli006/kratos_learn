# kratos学习

## protoc安装
- https://github.com/protocolbuffers/protobuf/releases 下载地址

## protoc-gen-go安装
- go get github.com/golang/protobuf/protoc-gen-go

## kratos安装
- go install github.com/go-kratos/kratos/cmd/kratos/v2@latest

## 创建项目
- cd ./backend
- kratos new verifyCode -r https://gitee.com/go-kratos/kratos-layout.git
- go mod tidy
- go get github.com/google/wire/cmd/wire
- go generate ./...

## 添加代码
- kratos proto add api/verifyCode/verifyCode.proto
- kratos proto client api\verifyCode\verifyCode.proto
- kratos proto server api\verifyCode\verifyCode.proto -t .\internal\server 

## 项目注册
- internal/service/service.go 注册 
- ProviderSet = wire.NewSet(NewGreeterService, NewVerifyCodeService) 构建函数
- internal/server/grpc 添加参数 完成注册
- verifyCodeService *service.VerifyCodeService 15行
- verifyCode.RegisterVerifyCodeServer(srv, verifyCodeService) 39 行
- 重新执行 go generate ./...

## 再调用服务中生成存根
- customer 中添加verifyCode.proto
- kratos proto client ./api/verifyCode/verifyCode.proto
- 


