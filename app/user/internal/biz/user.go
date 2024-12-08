package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// User is a User model.
type User struct {
	ID           int
	Name         string
	Email        string
	Avatar       string
	Professional string
	Skills       []string
	Description  string
	Address      string
	Experience   int
	Project      int
	CreatedAt    int64
	UpdatedAt    int64
}

// UserRepo is a Greater repo.
type UserRepo interface {
	SaveUser(context.Context, *User) error
	GetUser(context.Context, *User) (*User, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateUser creates a User, and returns the new User.
func (uc *UserUsecase) CreateUser(ctx context.Context, g *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
