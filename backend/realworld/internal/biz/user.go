package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Email        string
	UserName     string
	Token        string
	Bio          string
	Image        string
	PasswordHash string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
}
type ProfileRepo interface {
}

type UserUsecase struct {
	ur UserRepo
	pr ProfileRepo

	log *log.Helper
}

func NewUserUsecase(ur UserRepo, pr ProfileRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, pr: pr, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Register(ctx context.Context, u *User) error {
	err := uc.ur.CreateUser(ctx, u)
	if err != nil {

	}
	return nil
}
