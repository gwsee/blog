package biz

import (
	"blog/api/account/v1"
	"blog/app/account/internal/conf"
	"blog/internal/common"
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
		log:  log.NewHelper(log.With(logger, "module", "app/account-account-service")),
	}
}

func (uc *AccountUseCase) CreateAccount(ctx context.Context, req *v1.CreateAccountRequest) (reply *v1.CreateAccountReply, err error) {
	if req.Account == "" || req.Password == "" {
		return nil, errors.New("account or password is empty")
	}
	//两次MD5 //前端传过来应该已经MD5了  然后我们再次MD5才行
	req.Password = common.MD5(req.Password)
	err = uc.repo.CreateAccount(ctx, &Account{
		Account:  req.Account,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return
	}
	return &v1.CreateAccountReply{}, err
}
func (uc *AccountUseCase) ResetPassword(ctx context.Context, req *v1.ResetPasswordRequest) (reply *v1.ResetPasswordReply, err error) {
	return nil, nil
}
func (uc *AccountUseCase) LoginByAccount(ctx context.Context, req *v1.LoginByAccountRequest) (reply *v1.LoginByAccountReply, err error) {
	account, err := uc.repo.FindByValidAccount(ctx, req.Account)
	if err != nil {
		return
	}
	if account.Password != common.MD5(req.Password) {
		return nil, errors.New("account password error")
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": account.Id,
		"account": req.Account,
	})
	signedStr, err := claims.SignedString([]byte(uc.key))
	if err != nil {
		return nil, err
	}
	return &v1.LoginByAccountReply{
		Token: signedStr,
	}, nil
}
