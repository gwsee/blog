package data

import (
	"blog/api/global"
	v1 "blog/api/palaces/v1"
	"blog/app/bff/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type palacesRepo struct {
	data *Data
	log  *log.Helper
}

func NewPalacesRepo(data *Data, logger log.Logger) biz.PalaceRepo {
	return &palacesRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (o *palacesRepo) CreatePalaces(ctx context.Context, req *v1.CreatePalacesRequest) (*v1.CreatePalacesReply, error) {
	return o.data.pc.CreatePalaces(ctx, req)
}
func (o *palacesRepo) UpdatePalaces(ctx context.Context, req *v1.UpdatePalacesRequest) (*v1.UpdatePalacesReply, error) {
	return o.data.pc.UpdatePalaces(ctx, req)
}
func (o *palacesRepo) DeletePalaces(ctx context.Context, req *v1.DeletePalacesRequest) (*v1.DeletePalacesReply, error) {
	return o.data.pc.DeletePalaces(ctx, req)
}
func (o *palacesRepo) GetPalaces(ctx context.Context, req *v1.GetPalacesRequest) (*v1.GetPalacesReply, error) {
	return o.data.pc.GetPalaces(ctx, req)
}
func (o *palacesRepo) ListPalaces(ctx context.Context, req *v1.ListPalacesRequest) (*v1.ListPalacesReply, error) {
	return o.data.pc.ListPalaces(ctx, req)
}
func (o *palacesRepo) SaveMemo(ctx context.Context, req *v1.SaveMemoRequest) (*global.Empty, error) {
	return o.data.pc.SaveMemo(ctx, req)
}
func (o *palacesRepo) DoneMemo(ctx context.Context, req *global.State) (*global.Empty, error) {
	return o.data.pc.DoneMemo(ctx, req)
}
func (o *palacesRepo) DeleteMemo(ctx context.Context, req *global.ID) (*global.Empty, error) {
	return o.data.pc.DeleteMemo(ctx, req)
}
func (o *palacesRepo) ListMemo(ctx context.Context, req *v1.ListMemoRequest) (*v1.ListMemoReply, error) {
	return o.data.pc.ListMemo(ctx, req)
}
func (o *palacesRepo) SaveTodo(ctx context.Context, req *v1.SaveTodoItem) (*global.Empty, error) {
	return o.data.pc.SaveTodo(ctx, req)
}
func (o *palacesRepo) DoneTodo(ctx context.Context, id *global.ID) (*global.Empty, error) {
	return o.data.pc.DoneTodo(ctx, id)
}
func (o *palacesRepo) DeleteTodo(ctx context.Context, id *global.ID) (*global.Empty, error) {
	return o.data.pc.DeleteTodo(ctx, id)
}
func (o *palacesRepo) DeleteTodoDone(ctx context.Context, id *global.ID) (*global.Empty, error) {
	return o.data.pc.DeleteTodoDone(ctx, id)
}
func (o *palacesRepo) DeleteTodoRecord(ctx context.Context, id *global.ID) (*global.Empty, error) {
	return o.data.pc.DeleteTodoRecord(ctx, id)
}
func (o *palacesRepo) ListTodo(ctx context.Context, req *v1.ListTodoRequest) (*v1.ListTodoReply, error) {
	return o.data.pc.ListTodo(ctx, req)
}
func (o *palacesRepo) ListTodoDone(ctx context.Context, req *v1.ListTodoRequest) (*v1.ListTodoReply, error) {
	return o.data.pc.ListTodoDone(ctx, req)
}
