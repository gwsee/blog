package data

import (
	"blog/app/account/internal/biz"
	"blog/internal/ent/account"
	"blog/internal/ent/migrate"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

const (
	Invalid = iota
	Valid
)

type accountRepo struct {
	data *Data
	log  *log.Helper
}

func NewAccountRepo(data *Data, logger log.Logger) biz.AccountRepo {
	if err := data.db.Schema.Create(context.TODO(), migrate.WithForeignKeys(false),
		migrate.WithDropColumn(true)); err != nil {
		panic(err)
	}
	return &accountRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (o *accountRepo) CreateAccount(ctx context.Context, data *biz.Account) error {
	_, err := o.data.db.Account.Create().SetAccount(data.Account).SetPassword(data.Password).
		SetNickname(data.Nickname).SetAvatar(data.Avatar).SetDescription(data.Description).
		SetEmail(data.Email).SetStatus(Valid).Save(ctx)
	return err
}

func (o *accountRepo) ResetPassword(ctx context.Context, data *biz.Account) error {
	_, err := o.data.db.Account.UpdateOneID(data.Id).SetPassword(data.Password).Save(ctx)
	return err
}

func (o *accountRepo) FindByValidAccount(ctx context.Context, data string) (*biz.Account, error) {
	info, err := o.data.db.Account.Query().Where(account.StatusEQ(Valid)).
		Where(account.AccountEQ(data)).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Account{
		Id:       info.ID,
		Account:  info.Account,
		Password: info.Password,
	}, nil
}

func (o *accountRepo) Info(ctx context.Context, id int64) (*biz.Account, error) {
	info, err := o.data.db.Account.Query().Where(account.IDEQ(int(id))).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Account{
		Id:          info.ID,
		Account:     info.Account,
		Email:       info.Email,
		Nickname:    info.Nickname,
		Avatar:      info.Avatar,
		Description: info.Description,
		BlogNum:     info.BlogNum,
	}, err
}
func (o *accountRepo) UpdateAccount(ctx context.Context, data *biz.Account) error {
	_, err := o.data.db.Account.UpdateOneID(data.Id).SetAvatar(data.Avatar).
		SetNickname(data.Nickname).SetEmail(data.Email).SetDescription(data.Description).Save(ctx)
	return err
}
