package data

import (
	"context"
	"customer/api/verifyCode"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"time"
)

type CustomerData struct {
	data *Data
}

func NewCustomerData(data *Data) *CustomerData {
	return &CustomerData{data: data}
}

func (t CustomerData) SetVerifyCode(telephone, code string, ex int64) error {
	status := t.data.Rdb.Set(context.Background(), "CVC:"+telephone, code, time.Duration(ex)*time.Second)
	if _, err := status.Result(); err != nil {
		return err
	}
	return nil
}

func (t CustomerData) GetVerifyCode(telephone string) (string, error) {
	return t.data.Rdb.Get(context.Background(), "CVC:"+telephone).Result()
}

func (t CustomerData) CreateVerifyCode() (string, error) {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("127.0.0.1:9000"))
	if err != nil {
		return "", err
	}
	defer func() { _ = conn.Close() }()

	client := verifyCode.NewVerifyCodeClient(conn)

	reply, err := client.GetVerifyCode(context.Background(), &verifyCode.GetVerifyCodeRequest{
		Length: 6,
		Type:   3})
	if err != nil {
		return "", err
	}
	return reply.Code, nil
}
