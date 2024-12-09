package data

import (
	"blog/api/global"
	v1 "blog/api/user/v1"
	"blog/app/bff/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (o *userRepo) SaveUser(ctx context.Context, in *v1.SaveUserRequest) (*global.Empty, error) {
	return o.data.uc.SaveUser(ctx, in)
}
func (o *userRepo) GetUser(ctx context.Context, in *global.Empty) (*v1.GetUserReply, error) {
	return o.data.uc.GetUser(ctx, in)
}
func (o *userRepo) SaveProject(ctx context.Context, in *v1.SaveProjectRequest) (*global.Empty, error) {
	return o.data.uc.SaveProject(ctx, in)
}
func (o *userRepo) DeleteProject(ctx context.Context, in *global.ID) (*global.Empty, error) {
	return o.data.uc.DeleteProject(ctx, in)
}
func (o *userRepo) GetProject(ctx context.Context, in *global.ID) (*v1.GetProjectReply, error) {
	return o.data.uc.GetProject(ctx, in)
}
func (o *userRepo) ListProject(ctx context.Context, in *v1.ListProjectRequest) (*v1.ListProjectReply, error) {
	return o.data.uc.ListProject(ctx, in)
}
func (o *userRepo) SaveExperience(ctx context.Context, in *v1.SaveExperienceRequest) (*global.Empty, error) {
	return o.data.uc.SaveExperience(ctx, in)
}
func (o *userRepo) DeleteExperience(ctx context.Context, in *global.ID) (*global.Empty, error) {
	return o.data.uc.DeleteExperience(ctx, in)
}
func (o *userRepo) GetExperience(ctx context.Context, in *global.ID) (*v1.GetExperienceReply, error) {
	return o.data.uc.GetExperience(ctx, in)
}
func (o *userRepo) ListExperience(ctx context.Context, in *v1.ListExperienceRequest) (*v1.ListExperienceReply, error) {
	return o.data.uc.ListExperience(ctx, in)
}
