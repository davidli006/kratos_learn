package service

import (
	"context"
	v1 "realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (replay *v1.UserReplay, err error) {
	return &v1.UserReplay{User: &v1.UserReplay_User{
		Username: "kakaxi",
	}}, nil
}
