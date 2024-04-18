package service

import (
	"context"
	pb "realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReplay, error) {
	return &pb.UserReplay{User: &pb.UserReplay_User{
		Username: "kakaxi",
		Id:       1,
		IsDelete: 2,
	}}, nil
}

func (s *RealWorldService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReplay, error) {

	return &pb.UserReplay{}, nil
}

func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *pb.GetCurrentUserRequest) (*pb.UserReplay, error) {
	return &pb.UserReplay{}, nil
}

func (s *RealWorldService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserReplay, error) {
	return &pb.UserReplay{}, nil
}

func (s *RealWorldService) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.ProfileReplay, error) {
	return &pb.ProfileReplay{}, nil
}

func (s *RealWorldService) FollowUser(ctx context.Context, req *pb.FollowUserRequest) (*pb.ProfileReplay, error) {
	return &pb.ProfileReplay{}, nil
}

func (s *RealWorldService) UnFollowUser(ctx context.Context, req *pb.UnFollowUserRequest) (*pb.ProfileReplay, error) {
	return &pb.ProfileReplay{}, nil
}
