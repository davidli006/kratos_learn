package service

import (
	"github.com/google/wire"
	pb "realworld/api/realworld/v1"
	"realworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

type RealWorldService struct {
	pb.UnimplementedRealWorldServer

	uc *biz.UserUsecase
	su *biz.SocialUsecase
}

func NewRealWorldService(uc *biz.UserUsecase, su *biz.SocialUsecase) *RealWorldService {
	return &RealWorldService{uc: uc, su: su}
}
