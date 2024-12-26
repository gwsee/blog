package biz

import (
	"blog/api/global"
	v1 "blog/api/user/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}
type UserRepo interface {
	SaveUser(context.Context, *v1.SaveUserRequest) (*global.Empty, error)
	GetUser(context.Context, *global.Empty) (*v1.GetUserReply, error)
	SaveProject(context.Context, *v1.SaveProjectRequest) (*global.Empty, error)
	DeleteProject(context.Context, *global.ID) (*global.Empty, error)
	GetProject(context.Context, *global.ID) (*v1.GetProjectReply, error)
	ListProject(context.Context, *v1.ListProjectRequest) (*v1.ListProjectReply, error)
	SaveExperience(context.Context, *v1.SaveExperienceRequest) (*global.Empty, error)
	DeleteExperience(context.Context, *global.ID) (*global.Empty, error)
	GetExperience(context.Context, *global.ID) (*v1.GetExperienceReply, error)
	ListExperience(context.Context, *v1.ListExperienceRequest) (*v1.ListExperienceReply, error)
	Photos(context.Context, *v1.PhotosReq) (*v1.PhotosReply, error)
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}
func (uc *UserUseCase) SaveUser(ctx context.Context, req *v1.SaveUserRequest) (*global.Empty, error) {
	return uc.repo.SaveUser(ctx, req)
}
func (uc *UserUseCase) GetUser(ctx context.Context, req *global.Empty) (*v1.GetUserReply, error) {
	return uc.repo.GetUser(ctx, req)
}
func (uc *UserUseCase) SaveProject(ctx context.Context, req *v1.SaveProjectRequest) (*global.Empty, error) {
	return uc.repo.SaveProject(ctx, req)
}
func (uc *UserUseCase) DeleteProject(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return uc.repo.DeleteProject(ctx, req)
}
func (uc *UserUseCase) GetProject(ctx context.Context, req *global.ID) (*v1.GetProjectReply, error) {
	return uc.repo.GetProject(ctx, req)
}
func (uc *UserUseCase) ListProject(ctx context.Context, req *v1.ListProjectRequest) (*v1.ListProjectReply, error) {
	return uc.repo.ListProject(ctx, req)
}

func (uc *UserUseCase) SaveExperience(ctx context.Context, req *v1.SaveExperienceRequest) (*global.Empty, error) {
	return uc.repo.SaveExperience(ctx, req)
}
func (uc *UserUseCase) DeleteExperience(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return uc.repo.DeleteExperience(ctx, req)
}
func (uc *UserUseCase) GetExperience(ctx context.Context, req *global.ID) (*v1.GetExperienceReply, error) {
	return uc.repo.GetExperience(ctx, req)
}
func (uc *UserUseCase) ListExperience(ctx context.Context, req *v1.ListExperienceRequest) (*v1.ListExperienceReply, error) {
	return uc.repo.ListExperience(ctx, req)
}
func (uc *UserUseCase) Photos(ctx context.Context, req *v1.PhotosReq) (*v1.PhotosReply, error) {
	return uc.repo.Photos(ctx, req)
}
