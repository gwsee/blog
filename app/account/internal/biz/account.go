package biz

import (
	"blog/api/account/v1"
	"blog/api/global"
	"blog/internal/common"
	"blog/internal/constx"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

type Account struct {
	Id          int
	Account     string
	Password    string
	Email       string
	Nickname    string
	Avatar      string
	Description string
	BlogNum     int
}

// AccountRepo is a  repo.
type AccountRepo interface {
	CreateAccount(context.Context, *Account) error
	ResetPassword(context.Context, *Account) error
	FindByValidAccount(context.Context, string) (*Account, error)
	UpdateAccount(context.Context, *Account) error
	Info(context.Context, int64) (*Account, error)
}
type AccountUseCase struct {
	repo AccountRepo
	log  *log.Helper
	key  string
}

func NewAccountUseCase(conf *global.Auth, repo AccountRepo, logger log.Logger) *AccountUseCase {
	return &AccountUseCase{
		key:  conf.ApiKey,
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "app/account-account-service")),
	}
}

func (uc *AccountUseCase) CreateAccount(ctx context.Context, req *v1.CreateAccountRequest) (reply *global.Empty, err error) {
	if req.Account == "" || req.Password == "" {
		return nil, errors.New("account or password is empty")
	}
	//两次MD5 //前端传过来应该已经MD5了  然后我们再次MD5才行
	req.Password = common.MD5(req.Password)
	err = uc.repo.CreateAccount(ctx, &Account{
		Account:     req.Account,
		Password:    req.Password,
		Email:       req.Email,
		Nickname:    req.Nickname,
		Avatar:      req.Avatar,
		Description: req.Description,
	})
	if err != nil {
		return
	}
	return &global.Empty{}, err
}
func (uc *AccountUseCase) ResetPassword(ctx context.Context, req *v1.ResetPasswordRequest) (reply *global.Empty, err error) {
	user, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	if user.Account != req.Account {
		return &global.Empty{}, nil
	}
	req.Password = common.MD5(req.Password)
	err = uc.repo.ResetPassword(ctx, &Account{
		Id:       int(user.Id),
		Password: req.Password,
	})
	return &global.Empty{}, err
}
func (uc *AccountUseCase) LoginByAccount(ctx context.Context, req *v1.LoginByAccountRequest) (reply *v1.LoginByAccountReply, err error) {
	account, err := uc.repo.FindByValidAccount(ctx, req.Account)
	if err != nil {
		return
	}
	if account.Password != common.MD5(req.Password) {
		return nil, errors.New("account password error")
	}
	by, _ := json.Marshal(constx.User{
		Id:      int64(account.Id),
		Account: account.Account,
	})
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		constx.UserInfo: string(by),
	})
	signedStr, err := claims.SignedString([]byte(uc.key))
	if err != nil {
		return nil, err
	}
	return &v1.LoginByAccountReply{
		Token: signedStr,
	}, nil
}

func (uc *AccountUseCase) Info(ctx context.Context) (reply *v1.AccountInfoReply, err error) {
	user, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	info, err := uc.repo.Info(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	return &v1.AccountInfoReply{
		Id:          int64(info.Id),
		Nickname:    info.Nickname,
		Account:     info.Account,
		Email:       info.Email,
		Avatar:      info.Avatar,
		Description: info.Description,
		BlogNum:     int64(info.BlogNum),
	}, nil
}
func (uc *AccountUseCase) UpdateAccount(ctx context.Context, req *v1.UpdateAccountRequest) (reply *global.Empty, err error) {
	user, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	err = uc.repo.UpdateAccount(ctx, &Account{
		Id:          int(user.Id),
		Nickname:    req.Nickname,
		Email:       req.Email,
		Avatar:      req.Avatar,
		Description: req.Description,
	})
	return &global.Empty{}, err
}
