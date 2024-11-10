package biz

import (
	account "blog/api/account/v1"
	"blog/api/global"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/anypb"
)

type AccountUseCase struct {
	repo AccountRepo
	log  *log.Helper
}

type AccountRepo interface {
	CreateAccount(ctx context.Context, data *account.CreateAccountRequest) (*account.CreateAccountReply, error)
	ResetPassword(ctx context.Context, data *account.ResetPasswordRequest) (*account.ResetPasswordReply, error)
	LoginByAccount(ctx context.Context, data *account.LoginByAccountRequest) (*account.LoginByAccountReply, error)
}

func NewAccountUseCase(repo AccountRepo, logger log.Logger) *AccountUseCase {
	return &AccountUseCase{repo: repo, log: log.NewHelper(logger)}
}
func (l *AccountUseCase) CreateAccount(ctx context.Context, request *account.CreateAccountRequest) (*global.Response, error) {
	res, err := l.repo.CreateAccount(ctx, request)
	if err != nil {
		return nil, err
	}
	anyData, err := anypb.New(res)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}
	return &global.Response{
		Data: anyData,
	}, nil
}
func (l *AccountUseCase) ResetPassword(ctx context.Context, request *account.ResetPasswordRequest) (*global.Response, error) {
	res, err := l.repo.ResetPassword(ctx, request)
	if err != nil {
		return nil, err
	}
	anyData, err := anypb.New(res)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}
	return &global.Response{
		Data: anyData,
	}, nil
}
func (l *AccountUseCase) LoginByAccount(ctx context.Context, request *account.LoginByAccountRequest) (*global.Response, error) {
	res, err := l.repo.LoginByAccount(ctx, request)
	if err != nil {
		return nil, err
	}
	anyData, err := anypb.New(res)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}
	return &global.Response{
		Data: anyData,
	}, nil
}
