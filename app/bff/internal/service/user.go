package service

import (
	pb "blog/api/bff/v1"
	"blog/api/global"
	v1 "blog/api/user/v1"
	"blog/app/bff/internal/biz"
	"context"
)

type UserService struct {
	pb.UnimplementedUserServer
	repo *biz.UserUseCase
}

func NewUserService(biz *biz.UserUseCase) *UserService {
	return &UserService{repo: biz}
}
func (uc *UserService) SaveUser(ctx context.Context, req *v1.SaveUserRequest) (*global.Empty, error) {
	return uc.repo.SaveUser(ctx, req)
}
func (uc *UserService) GetUser(ctx context.Context, req *global.Empty) (*v1.GetUserReply, error) {
	return uc.repo.GetUser(ctx, req)
}
func (uc *UserService) SaveProject(ctx context.Context, req *v1.SaveProjectRequest) (*global.Empty, error) {
	return uc.repo.SaveProject(ctx, req)
}
func (uc *UserService) DeleteProject(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return uc.repo.DeleteProject(ctx, req)
}
func (uc *UserService) GetProject(ctx context.Context, req *global.ID) (*v1.GetProjectReply, error) {
	return uc.repo.GetProject(ctx, req)
}
func (uc *UserService) ListProject(ctx context.Context, req *v1.ListProjectRequest) (*v1.ListProjectReply, error) {
	return uc.repo.ListProject(ctx, req)
}

func (uc *UserService) SaveExperience(ctx context.Context, req *v1.SaveExperienceRequest) (*global.Empty, error) {
	return uc.repo.SaveExperience(ctx, req)
}
func (uc *UserService) DeleteExperience(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return uc.repo.DeleteExperience(ctx, req)
}
func (uc *UserService) GetExperience(ctx context.Context, req *global.ID) (*v1.GetExperienceReply, error) {
	return uc.repo.GetExperience(ctx, req)
}
func (uc *UserService) ListExperience(ctx context.Context, req *v1.ListExperienceRequest) (*v1.ListExperienceReply, error) {
	return uc.repo.ListExperience(ctx, req)
}
func (uc *UserService) Photos(ctx context.Context, req *v1.PhotosReq) (*v1.PhotosReply, error) {
	return uc.repo.Photos(ctx, req)
}
