package biz

import (
	account "blog/api/account/v1"
	"blog/api/global"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type AccountUseCase struct {
	repo AccountRepo
	log  *log.Helper
}

type AccountRepo interface {
	CreateAccount(ctx context.Context, data *account.CreateAccountRequest) (*global.Empty, error)
	ResetPassword(ctx context.Context, data *account.ResetPasswordRequest) (*global.Empty, error)
	LoginByAccount(ctx context.Context, data *account.LoginByAccountRequest) (*account.LoginByAccountReply, error)
	Info(ctx context.Context, data *global.Empty) (*account.AccountInfoReply, error)
	UpdateAccount(ctx context.Context, data *account.UpdateAccountRequest) (*global.Empty, error)
}

func NewAccountUseCase(repo AccountRepo, logger log.Logger) *AccountUseCase {
	return &AccountUseCase{repo: repo, log: log.NewHelper(logger)}
}
func (l *AccountUseCase) CreateAccount(ctx context.Context, request *account.CreateAccountRequest) (*global.Empty, error) {
	return l.repo.CreateAccount(ctx, request)
	//if err != nil {
	//	return nil, err
	//}
	//anyData, err := anypb.New(res)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal response data: %w", err)
	//}
	//return &global.Response{
	//	Data: anyData,
	//}, nil
}
func (l *AccountUseCase) ResetPassword(ctx context.Context, request *account.ResetPasswordRequest) (*global.Empty, error) {
	return l.repo.ResetPassword(ctx, request)
	//if err != nil {
	//	return nil, err
	//}
	//anyData, err := anypb.New(res)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal response data: %w", err)
	//}
	//return &global.Response{
	//	Data: anyData,
	//}, nil
}
func (l *AccountUseCase) LoginByAccount(ctx context.Context, request *account.LoginByAccountRequest) (*account.LoginByAccountReply, error) {
	return l.repo.LoginByAccount(ctx, request)
	//if err != nil {
	//	return nil, err
	//}
	//anyData, err := anypb.New(res)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal response data: %w", err)
	//}
	//return &global.Response{
	//	Data: anyData,
	//}, nil
}
func (l *AccountUseCase) Info(ctx context.Context, data *global.Empty) (*account.AccountInfoReply, error) {
	return l.repo.Info(ctx, data)
	//if err != nil {
	//	return nil, err
	//}
	//anyData, err := anypb.New(res)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal response data: %w", err)
	//}
	//return &global.Response{
	//	Data: anyData,
	//}, nil
}
func (l *AccountUseCase) UpdateAccount(ctx context.Context, data *account.UpdateAccountRequest) (*global.Empty, error) {
	return l.repo.UpdateAccount(ctx, data)
	//if err != nil {
	//	return nil, err
	//}
	//anyData, err := anypb.New(res)
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal response data: %w", err)
	//}
	//return &global.Response{
	//	Data: anyData,
	//}, nil
}
