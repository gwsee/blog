package data

import (
	account "blog/api/account/v1"
	"blog/api/global"
	"blog/app/bff/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	return &accountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (l *accountRepo) CreateAccount(ctx context.Context, data *account.CreateAccountRequest) (*global.Empty, error) {
	return l.data.ac.CreateAccount(ctx, data)
}
func (l *accountRepo) ResetPassword(ctx context.Context, data *account.ResetPasswordRequest) (*global.Empty, error) {
	return l.data.ac.ResetPassword(ctx, data)
}
func (l *accountRepo) LoginByAccount(ctx context.Context, data *account.LoginByAccountRequest) (*account.LoginByAccountReply, error) {
	return l.data.ac.LoginByAccount(ctx, data)
}
func (l *accountRepo) Info(ctx context.Context, data *global.Empty) (*account.AccountInfoReply, error) {
	return l.data.ac.Info(ctx, data)
}
func (l *accountRepo) UpdateAccount(ctx context.Context, data *account.UpdateAccountRequest) (*global.Empty, error) {
	return l.data.ac.UpdateAccount(ctx, data)
}
