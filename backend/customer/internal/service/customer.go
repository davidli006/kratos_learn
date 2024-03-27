package service

import (
	"context"
	"customer/internal/data"
	"regexp"
	"time"

	pb "customer/api/customer"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
	Cd *data.CustomerData
}

func NewCustomerService(cd *data.CustomerData) *CustomerService {
	return &CustomerService{
		Cd: cd,
	}
}

func (s *CustomerService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeReq) (*pb.GetVerifyCodeResp, error) {
	telephone := req.Telephone
	pattern := `^1\d{10}$|^(0\d{2,3}-?|\(0\d{2,3}\))?[1-9]\d{4,7}(-\d{1,8})?$`
	regexPattern := regexp.MustCompile(pattern)
	if !regexPattern.MatchString(telephone) {
		return &pb.GetVerifyCodeResp{
			Code:    100,
			Message: "手机号码有误!"}, nil
	}

	// 获取验证码 服务
	code, err := s.Cd.CreateVerifyCode()
	if err != nil {
		return &pb.GetVerifyCodeResp{
			Code:    100,
			Message: "验证码创建失败!",
		}, nil
	}

	// redis存储
	err = s.Cd.SetVerifyCode(telephone, code, 60)
	if err != nil {
		return &pb.GetVerifyCodeResp{
			Code:    100,
			Message: "验证码保存失败!",
		}, nil
	}

	return &pb.GetVerifyCodeResp{
		Code:           0,
		VerifyCode:     code,
		VerifyCodeTime: time.Now().Unix(),
		VerifyCodeLife: 60,
	}, nil
}
