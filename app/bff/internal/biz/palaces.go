package biz

import (
	"blog/api/global"
	v1 "blog/api/palaces/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type PalacesUseCase struct {
	repo PalaceRepo
	log  *log.Helper
}
type PalaceRepo interface {
	CreatePalaces(context.Context, *v1.CreatePalacesRequest) (*v1.CreatePalacesReply, error)
	UpdatePalaces(context.Context, *v1.UpdatePalacesRequest) (*v1.UpdatePalacesReply, error)
	DeletePalaces(context.Context, *v1.DeletePalacesRequest) (*v1.DeletePalacesReply, error)
	GetPalaces(context.Context, *v1.GetPalacesRequest) (*v1.GetPalacesReply, error)
	ListPalaces(context.Context, *v1.ListPalacesRequest) (*v1.ListPalacesReply, error)

	SaveMemo(context.Context, *v1.SaveMemoRequest) (*global.Empty, error)
	DoneMemo(context.Context, *global.State) (*global.Empty, error)
	DeleteMemo(context.Context, *global.ID) (*global.Empty, error)
	ListMemo(context.Context, *v1.ListMemoRequest) (*v1.ListMemoReply, error)

	SaveTodo(context.Context, *v1.SaveTodoItem) (*global.Empty, error)
	DoneTodo(context.Context, *global.ID) (*global.Empty, error)
	DeleteTodo(context.Context, *global.ID) (*global.Empty, error)
	DeleteTodoDone(context.Context, *global.ID) (*global.Empty, error)
	DeleteTodoRecord(context.Context, *global.ID) (*global.Empty, error)
	ListTodo(context.Context, *v1.ListTodoRequest) (*v1.ListTodoReply, error)
	ListTodoDone(context.Context, *v1.ListTodoRequest) (*v1.ListTodoReply, error)
}

func NewPalacesUseCase(repo PalaceRepo, logger log.Logger) *PalacesUseCase {
	return &PalacesUseCase{repo: repo, log: log.NewHelper(logger)}
}
func (o *PalacesUseCase) CreatePalaces(ctx context.Context, req *v1.CreatePalacesRequest) (*v1.CreatePalacesReply, error) {
	return o.repo.CreatePalaces(ctx, req)
}
func (o *PalacesUseCase) UpdatePalaces(ctx context.Context, req *v1.UpdatePalacesRequest) (*v1.UpdatePalacesReply, error) {
	return o.repo.UpdatePalaces(ctx, req)
}
func (o *PalacesUseCase) DeletePalaces(ctx context.Context, req *v1.DeletePalacesRequest) (*v1.DeletePalacesReply, error) {
	return o.repo.DeletePalaces(ctx, req)
}
func (o *PalacesUseCase) GetPalaces(ctx context.Context, req *v1.GetPalacesRequest) (*v1.GetPalacesReply, error) {
	return o.repo.GetPalaces(ctx, req)
}
func (o *PalacesUseCase) ListPalaces(ctx context.Context, req *v1.ListPalacesRequest) (*v1.ListPalacesReply, error) {
	return o.repo.ListPalaces(ctx, req)
}
func (o *PalacesUseCase) SaveMemo(ctx context.Context, req *v1.SaveMemoRequest) (*global.Empty, error) {
	return o.repo.SaveMemo(ctx, req)
}
func (o *PalacesUseCase) DoneMemo(ctx context.Context, req *global.State) (*global.Empty, error) {
	return o.repo.DoneMemo(ctx, req)
}
func (o *PalacesUseCase) DeleteMemo(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return o.repo.DeleteMemo(ctx, req)
}
func (o *PalacesUseCase) ListMemo(ctx context.Context, req *v1.ListMemoRequest) (*v1.ListMemoReply, error) {
	return o.repo.ListMemo(ctx, req)
}
func (o *PalacesUseCase) SaveTodo(ctx context.Context, req *v1.SaveTodoItem) (*global.Empty, error) {
	return o.repo.SaveTodo(ctx, req)
}
func (o *PalacesUseCase) DoneTodo(ctx context.Context, id *global.ID) (*global.Empty, error) {
	return o.repo.DoneTodo(ctx, id)
}
func (o *PalacesUseCase) DeleteTodo(ctx context.Context, id *global.ID) (*global.Empty, error) {
	return o.repo.DeleteTodo(ctx, id)
}
func (o *PalacesUseCase) DeleteTodoDone(ctx context.Context, id *global.ID) (*global.Empty, error) {
	return o.repo.DeleteTodoDone(ctx, id)
}
func (o *PalacesUseCase) DeleteTodoRecord(ctx context.Context, id *global.ID) (*global.Empty, error) {
	return o.repo.DeleteTodoRecord(ctx, id)
}
func (o *PalacesUseCase) ListTodo(ctx context.Context, req *v1.ListTodoRequest) (*v1.ListTodoReply, error) {
	return o.repo.ListTodo(ctx, req)
}
func (o *PalacesUseCase) ListTodoDone(ctx context.Context, req *v1.ListTodoRequest) (*v1.ListTodoReply, error) {
	return o.repo.ListTodoDone(ctx, req)
}
