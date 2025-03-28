package service

import (
	"blog/api/global"
	"blog/app/bff/internal/biz"
	"context"

	pb "blog/api/bff/v1"
	palaces "blog/api/palaces/v1"
)

type PalacesService struct {
	pb.UnimplementedPalacesServer
	uc *biz.PalacesUseCase
}

func NewPalacesService(uc *biz.PalacesUseCase) *PalacesService {
	return &PalacesService{uc: uc}
}

func (s *PalacesService) CreatePalaces(ctx context.Context, req *palaces.CreatePalacesRequest) (*palaces.CreatePalacesReply, error) {
	return s.uc.CreatePalaces(ctx, req)
}
func (s *PalacesService) UpdatePalaces(ctx context.Context, req *palaces.UpdatePalacesRequest) (*palaces.UpdatePalacesReply, error) {
	return s.uc.UpdatePalaces(ctx, req)
}
func (s *PalacesService) DeletePalaces(ctx context.Context, req *palaces.DeletePalacesRequest) (*palaces.DeletePalacesReply, error) {
	return s.uc.DeletePalaces(ctx, req)
}
func (s *PalacesService) GetPalaces(ctx context.Context, req *palaces.GetPalacesRequest) (*palaces.GetPalacesReply, error) {
	return s.uc.GetPalaces(ctx, req)
}
func (s *PalacesService) ListPalaces(ctx context.Context, req *palaces.ListPalacesRequest) (*palaces.ListPalacesReply, error) {
	return s.uc.ListPalaces(ctx, req)
}
func (s *PalacesService) SaveMemo(ctx context.Context, req *palaces.SaveMemoRequest) (*global.Empty, error) {
	return s.uc.SaveMemo(ctx, req)
}
func (s *PalacesService) DoneMemo(ctx context.Context, req *global.State) (*global.Empty, error) {
	return s.uc.DoneMemo(ctx, req)
}
func (s *PalacesService) DeleteMemo(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return s.uc.DeleteMemo(ctx, req)
}
func (s *PalacesService) ListMemo(ctx context.Context, req *palaces.ListMemoRequest) (*palaces.ListMemoReply, error) {
	return s.uc.ListMemo(ctx, req)
}
func (s *PalacesService) SaveTodo(ctx context.Context, req *palaces.SaveTodoItem) (*global.Empty, error) {
	return s.uc.SaveTodo(ctx, req)
}
func (s *PalacesService) DoneTodo(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return s.uc.DoneTodo(ctx, req)
}
func (s *PalacesService) DeleteTodo(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return s.uc.DeleteTodo(ctx, req)
}
func (s *PalacesService) DeleteTodoDone(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return s.uc.DeleteTodoDone(ctx, req)
}
func (s *PalacesService) DeleteTodoRecord(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return s.uc.DeleteTodoRecord(ctx, req)
}
func (s *PalacesService) ListTodo(ctx context.Context, req *palaces.ListTodoRequest) (*palaces.ListTodoReply, error) {
	return s.uc.ListTodo(ctx, req)
}
func (s *PalacesService) ListTodoDone(ctx context.Context, req *palaces.ListTodoRequest) (*palaces.ListTodoReply, error) {
	return s.uc.ListTodoDone(ctx, req)
}
