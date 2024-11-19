package service

import (
	"blog/api/global"
	"blog/app/account/internal/biz"
	"context"

	pb "blog/api/account/v1"
)

type AccountService struct {
	pb.UnimplementedAccountServer
	ac *biz.AccountUseCase
}

func NewAccountService(ac *biz.AccountUseCase) *AccountService {
	return &AccountService{ac: ac}
}

func (s *AccountService) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*global.Empty, error) {
	return s.ac.CreateAccount(ctx, req)
}
func (s *AccountService) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*global.Empty, error) {
	return s.ac.ResetPassword(ctx, req)
}
func (s *AccountService) LoginByAccount(ctx context.Context, req *pb.LoginByAccountRequest) (*pb.LoginByAccountReply, error) {
	return s.ac.LoginByAccount(ctx, req)
}
func (s *AccountService) Info(ctx context.Context, req *global.Empty) (*pb.AccountInfoReply, error) {
	return s.ac.Info(ctx)
}
func (s *AccountService) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*global.Empty, error) {
	return s.ac.UpdateAccount(ctx, req)
}
