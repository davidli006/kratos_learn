package service

import (
	"context"
	"customer/api/verifyCode"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"regexp"
	"time"

	pb "customer/api/customer"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeReq) (*pb.GetVerifyCodeResp, error) {
	telephone := req.Telephone
	fmt.Printf("%s\n", telephone)
	pattern := `^1\d{10}$|^(0\d{2,3}-?|\(0\d{2,3}\))?[1-9]\d{4,7}(-\d{1,8})?$`
	regexPattern := regexp.MustCompile(pattern)
	if !regexPattern.MatchString(telephone) {
		return &pb.GetVerifyCodeResp{
			Code:    100,
			Message: "手机号码有误!"}, nil
	}

	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("127.0.0.1:9000"))
	if err != nil {
		return &pb.GetVerifyCodeResp{Code: 100, Message: "GRPC链接失败!"}, nil
	}
	defer func() { _ = conn.Close() }()

	client := verifyCode.NewVerifyCodeClient(conn)

	reply, err := client.GetVerifyCode(context.Background(), &verifyCode.GetVerifyCodeRequest{
		Length: 6,
		Type:   3,
	})
	if err != nil {
		return &pb.GetVerifyCodeResp{Code: 1, Message: "验证码获取失败!"}, nil
	}

	return &pb.GetVerifyCodeResp{
		Code:           0,
		VerifyCode:     reply.Code,
		VerifyCodeTime: time.Now().Unix(),
		VerifyCodeLife: 60,
	}, nil
}
