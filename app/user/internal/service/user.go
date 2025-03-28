package service

import (
	"blog/api/global"
	"blog/api/user/v1"
	"blog/app/user/internal/biz"
	"context"
)

// UserService is a greeter service.
type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUseCase
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}
func (s *UserService) SaveUser(ctx context.Context, req *v1.SaveUserRequest) (*global.Empty, error) {
	return s.uc.SaveUser(ctx, req)
}
func (s *UserService) GetUser(ctx context.Context, req *global.Empty) (*v1.GetUserReply, error) {
	return s.uc.GetUser(ctx)
}
func (s *UserService) SaveProject(ctx context.Context, req *v1.SaveProjectRequest) (*global.Empty, error) {
	return s.uc.SaveProject(ctx, req)
}
func (s *UserService) DeleteProject(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return s.uc.DeleteProject(ctx, req)
}
func (s *UserService) GetProject(ctx context.Context, req *global.ID) (*v1.GetProjectReply, error) {
	return s.uc.GetProject(ctx, req)
}
func (s *UserService) ListProject(ctx context.Context, req *v1.ListProjectRequest) (*v1.ListProjectReply, error) {
	return s.uc.ListProject(ctx, req)
}
func (s *UserService) SaveExperience(ctx context.Context, req *v1.SaveExperienceRequest) (*global.Empty, error) {
	return s.uc.SaveExperience(ctx, req)
}
func (s *UserService) DeleteExperience(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return s.uc.DeleteExperience(ctx, req)
}
func (s *UserService) GetExperience(ctx context.Context, req *global.ID) (*v1.GetExperienceReply, error) {
	return s.uc.GetExperience(ctx, req)
}
func (s *UserService) ListExperience(ctx context.Context, req *v1.ListExperienceRequest) (*v1.ListExperienceReply, error) {
	return s.uc.ListExperience(ctx, req)
}
func (s *UserService) Photos(ctx context.Context, req *v1.PhotosReq) (*v1.PhotosReply, error) {
	return s.uc.Photos(ctx, req)
}
func (s *UserService) Messages(ctx context.Context, req *global.PageInfo) (*v1.MessagesReply, error) {
	return s.uc.Messages(ctx, req)
}
