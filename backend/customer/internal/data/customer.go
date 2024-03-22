package data

import (
	"context"
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
