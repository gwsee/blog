package biz

import (
	"blog/api/global"
	v1 "blog/api/user/v1"
	"blog/internal/constx"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"math/rand"
	"time"
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
	*global.PageInfo
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
	*global.PageInfo
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
type PhotosQuery struct {
	PageSize int64
	Me       bool
	Type     string
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

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) SaveUser(ctx context.Context, req *v1.SaveUserRequest) (*global.Empty, error) {
	if req.Name == "" {
		return nil, errors.New("姓名必须")
	}
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	return &global.Empty{}, uc.repo.SaveUser(ctx, &User{
		ID:           int(u.Id),
		Name:         req.Name,
		Email:        req.Email,
		Avatar:       req.Avatar,
		Professional: req.Professional,
		Skills:       req.Skills,
		Description:  req.Description,
		Address:      req.Address,
	})
}
func (uc *UserUsecase) GetUser(ctx context.Context) (*v1.GetUserReply, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	info, err := uc.repo.GetUser(ctx, &User{ID: int(u.Id)})
	if err != nil {
		return nil, err
	}
	return &v1.GetUserReply{
		Id:           int64(info.ID),
		Name:         info.Name,
		Email:        info.Email,
		Avatar:       info.Avatar,
		Professional: info.Professional,
		Skills:       info.Skills,
		Description:  info.Description,
		Address:      info.Address,
		Experience:   int64(info.Experience),
		Project:      int64(info.Project),
		CreatedAt:    info.CreatedAt,
		UpdatedAt:    info.UpdatedAt,
	}, nil
}
func (uc *UserUsecase) SaveProject(ctx context.Context, req *v1.SaveProjectRequest) (*global.Empty, error) {
	if req.ExperienceId == 0 || req.Start == 0 || len(req.Title) == 0 {
		return nil, errors.New("参数丢失")
	}
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	return &global.Empty{}, uc.repo.SaveProject(ctx, &Project{
		ID:           int(req.Id),
		UserId:       int(u.Id),
		ExperienceId: int(req.ExperienceId),
		Title:        req.Title,
		Description:  req.Description,
		Skills:       req.Skills,
		Start:        req.Start,
		End:          req.End,
		Link:         req.Link,
		Photos:       req.Photos,
	})
}
func (uc *UserUsecase) GetProject(ctx context.Context, req *global.ID) (*v1.GetProjectReply, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	info, err := uc.repo.GetProject(ctx, &Project{ID: int(req.Id), UserId: int(u.Id)})
	if err != nil {
		return nil, err
	}
	return &v1.GetProjectReply{
		Id:           int64(info.ID),
		ExperienceId: int64(info.ExperienceId),
		Title:        info.Title,
		Description:  info.Description,
		Skills:       info.Skills,
		Start:        info.Start,
		End:          info.End,
		Link:         info.Link,
		Photos:       info.Photos,
		CreatedAt:    info.CreatedAt,
		UpdatedAt:    info.UpdatedAt,
		UserId:       int64(info.UserId),
	}, nil
}
func (uc *UserUsecase) DeleteProject(ctx context.Context, req *global.ID) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	return &global.Empty{}, uc.repo.DeleteProject(ctx, &Project{ID: int(req.Id), UserId: int(u.Id)})
}
func (uc *UserUsecase) ListProject(ctx context.Context, req *v1.ListProjectRequest) (*v1.ListProjectReply, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	total, list, err := uc.repo.ListProject(ctx, &ProjectQuery{
		PageInfo:     req.Page,
		ExperienceId: int(req.ExperienceId),
		UserId:       int(u.Id),
		Title:        req.Title,
		Skills:       req.Skills,
		Time:         req.Time,
	})
	if err != nil {
		return nil, err
	}
	resp := &v1.ListProjectReply{
		List:  make([]*v1.ListProject, 0, len(list)),
		Total: int64(total),
	}
	_rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, info := range list {
		one := &v1.ListProject{
			Id:           int64(info.ID),
			ExperienceId: int64(info.ExperienceId),
			Title:        info.Title,
			Skills:       info.Skills,
			Start:        info.Start,
			End:          info.End,
			Link:         info.Link,
			UpdatedAt:    info.UpdatedAt,
			UserId:       int64(info.UserId),
		}
		if len(info.Photos) > 0 {
			index := _rand.Intn(len(info.Photos))
			one.Photo = info.Photos[index]
		}
		resp.List = append(resp.List, one)
	}
	return resp, nil
}
func (uc *UserUsecase) SaveExperience(ctx context.Context, req *v1.SaveExperienceRequest) (*global.Empty, error) {
	if req.Start == 0 || len(req.Company) == 0 {
		return nil, errors.New("参数丢失")
	}
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	return &global.Empty{}, uc.repo.SaveExperience(ctx, &Experience{
		ID:               int(req.Id),
		UserId:           int(u.Id),
		Company:          req.Company,
		Role:             req.Role,
		Location:         req.Location,
		Start:            req.Start,
		End:              req.End,
		Description:      req.Description,
		Responsibilities: req.Responsibilities,
		Achievements:     req.Achievements,
		Skills:           req.Skills,
	})
}
func (uc *UserUsecase) GetExperience(ctx context.Context, req *global.ID) (*v1.GetExperienceReply, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	info, err := uc.repo.GetExperience(ctx, &Experience{ID: int(req.Id), UserId: int(u.Id)})
	if err != nil {
		return nil, err
	}
	return &v1.GetExperienceReply{
		Id:               int64(info.ID),
		Company:          info.Company,
		Role:             info.Role,
		Location:         info.Location,
		Start:            info.Start,
		End:              info.End,
		Description:      info.Description,
		Responsibilities: info.Responsibilities,
		Achievements:     info.Achievements,
		Skills:           info.Skills,
		CreatedAt:        info.CreatedAt,
		UpdatedAt:        info.UpdatedAt,
		UserId:           int64(info.UserId),
	}, nil
}
func (uc *UserUsecase) DeleteExperience(ctx context.Context, req *global.ID) (*global.Empty, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	return &global.Empty{}, uc.repo.DeleteExperience(ctx, &Experience{ID: int(req.Id), UserId: int(u.Id)})
}
func (uc *UserUsecase) ListExperience(ctx context.Context, req *v1.ListExperienceRequest) (*v1.ListExperienceReply, error) {
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	total, list, err := uc.repo.ListExperience(ctx, &ExperienceQuery{
		PageInfo: req.Page,
		UserId:   int(u.Id),
		Company:  req.Company,
	})
	if err != nil {
		return nil, err
	}
	resp := &v1.ListExperienceReply{
		Total: int64(total),
		List:  make([]*v1.ListExperience, 0, len(list)),
	}
	for _, info := range list {
		resp.List = append(resp.List, &v1.ListExperience{
			Id:          int64(info.ID),
			Company:     info.Company,
			Role:        info.Role,
			Location:    info.Location,
			Start:       info.Start,
			End:         info.End,
			Description: info.Description,
			UpdatedAt:   info.UpdatedAt,
			UserId:      int64(info.UserId),
		})
	}
	return resp, nil
}
