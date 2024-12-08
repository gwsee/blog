package biz

import (
	"blog/api/global"
	v1 "blog/api/user/v1"
	"blog/internal/constx"
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
type ProjectQuery struct {
	global.PageInfo
	ExperienceId int
	UserId       int
	Title        string
	Skills       []string
	Time         []int64
}
type Project struct {
	ID           int
	UserId       int
	CreatedAt    int64
	UpdatedAt    int64
	ExperienceId int
	Title        string
	Description  string
	Skills       []string
	Start        int64
	End          int64
	Link         string
	Photos       []string
}
type ExperienceQuery struct {
	global.PageInfo
	UserId  int
	Company string
}
type Experience struct {
	ID               int
	UserId           int
	CreatedAt        int64
	UpdatedAt        int64
	Company          string
	Role             string
	Location         string
	Start            int64
	End              int64
	Description      string
	Responsibilities string
	Achievements     string
	Skills           []string
}

// UserRepo is a Greater repo.
type UserRepo interface {
	SaveUser(context.Context, *User) error
	GetUser(context.Context, *User) (*User, error)

	SaveProject(context.Context, *Project) error
	GetProject(context.Context, *Project) (*Project, error)
	DeleteProject(context.Context, *Project) error
	ListProject(context.Context, *ProjectQuery) (int, []*Project, error)

	SaveExperience(context.Context, *Experience) error
	GetExperience(context.Context, *Experience) (*Experience, error)
	DeleteExperience(context.Context, *Experience) error
	ListExperience(context.Context, *ExperienceQuery) (int, []*Experience, error)
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
func (uc *UserUsecase) CreateUser(ctx context.Context, req *v1.SaveUserRequest) error {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return err
	}
	return uc.repo.SaveUser(ctx)
}
func (uc *UserUsecase) GetUser(ctx context.Context, g *User) (*User, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	return uc.repo.GetUser(ctx, g)
}
func (uc *UserUsecase) SaveProject(ctx context.Context, g *Project) error {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return err
	}
	return uc.repo.SaveProject(ctx, g)
}
func (uc *UserUsecase) GetProject(ctx context.Context, g *Project) (*Project, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	return uc.repo.GetProject(ctx, g)
}
func (uc *UserUsecase) DeleteProject(ctx context.Context, g *Project) error {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return err
	}
	return uc.repo.DeleteProject(ctx, g)
}
func (uc *UserUsecase) ListProject(ctx context.Context, g *ProjectQuery) (int, []*Project, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return
	}
	return uc.repo.ListProject(ctx, g)
}
func (uc *UserUsecase) SaveExperience(ctx context.Context, g *Experience) error {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return err
	}
	return uc.repo.SaveExperience(ctx, g)
}
func (uc *UserUsecase) GetExperience(ctx context.Context, g *Experience) (*Experience, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return
	}
	return uc.repo.GetExperience(ctx, g)
}
func (uc *UserUsecase) DeleteExperience(ctx context.Context, g *Experience) error {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return err
	}
	return uc.repo.DeleteExperience(ctx, g)
}
func (uc *UserUsecase) ListExperience(ctx context.Context, g *ExperienceQuery) (int, []*Experience, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return
	}
	return uc.repo.ListExperience(ctx, g)
}
