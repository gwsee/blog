package service

import (
	account "blog/api/account/v1"
	pb "blog/api/bff/v1"
	"blog/api/global"
	"blog/app/bff/internal/biz"
	"context"
)

type AccountService struct {
	pb.UnimplementedAccountServer
	biz *biz.AccountUseCase
}

func NewAccountService(biz *biz.AccountUseCase) *AccountService {
	return &AccountService{biz: biz}
}

func (s *AccountService) CreateAccount(ctx context.Context, req *account.CreateAccountRequest) (*global.Response, error) {
	return s.biz.CreateAccount(ctx, req)
}
func (s *AccountService) ResetPassword(ctx context.Context, req *account.ResetPasswordRequest) (*global.Response, error) {
	return s.biz.ResetPassword(ctx, req)
}
func (s *AccountService) LoginByAccount(ctx context.Context, req *account.LoginByAccountRequest) (*global.Response, error) {
	return s.biz.LoginByAccount(ctx, req)
}
func (s *AccountService) Info(ctx context.Context, req *global.Empty) (*global.Response, error) {
	return s.biz.Info(ctx, req)
}
func (s *AccountService) UpdateAccount(ctx context.Context, req *account.UpdateAccountRequest) (*global.Response, error) {
	return s.biz.UpdateAccount(ctx, req)
}
