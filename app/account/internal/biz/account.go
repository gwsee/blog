package biz

import (
	"blog/app/account/api"
	"blog/app/account/internal/conf"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

type Account struct {
	Id       int
	Account  string
	Password string
	Email    string
}

// AccountRepo is a  repo.
type AccountRepo interface {
	CreateAccount(context.Context, *Account) error
	ResetPassword(context.Context, *Account) error
	FindByValidAccount(context.Context, string) (*Account, error)
}
type AccountUseCase struct {
	repo AccountRepo
	log  *log.Helper
	key  string
}

func NewAccountUseCase(conf *conf.Auth, repo AccountRepo, logger log.Logger) *AccountUseCase {
	return &AccountUseCase{
		key:  conf.ApiKey,
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "app/account-service")),
	}
}

func (uc *AccountUseCase) CreateAccount(ctx context.Context, req *api.CreateAccountRequest) (reply *api.CreateAccountReply, err error) {
	if req.Account == "" || req.Password == "" {
		return nil, errors.New("account or password is empty")
	}
	err = uc.repo.CreateAccount(ctx, &Account{
		Account:  req.Account,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return
	}
	return &api.CreateAccountReply{}, err
}
func (uc *AccountUseCase) ResetPassword(ctx context.Context, req *api.ResetPasswordRequest) (reply *api.ResetPasswordReply, err error) {
	return nil, nil
}
func (uc *AccountUseCase) LoginByAccount(ctx context.Context, req *api.LoginByAccountRequest) (reply *api.LoginByAccountReply, err error) {
	account, err := uc.repo.FindByValidAccount(ctx, req.Account)
	if err != nil {
		return
	}
	if account.Password != req.Password {
		return nil, errors.New("account password error")
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": account.Id,
		"account": req.Account,
	})
	//TODO 设置key
	signedStr, err := claims.SignedString(uc.key)
	if err != nil {
		return nil, err
	}
	return &api.LoginByAccountReply{
		Token: signedStr,
	}, nil
}
