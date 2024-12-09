package biz

import (
	"blog/api/global"
	v1 "blog/api/tools/v1"
	"blog/internal/common"
	"blog/internal/constx"
	"context"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
)

// Tools is a Tools model.
type Tools struct {
	Hello string
}
type File struct {
	ID   string
	Type string
	Size int64
	Name string
	Path string
}

// ToolsRepo is a Greater repo.
type ToolsRepo interface {
	SaveFile(context.Context, *File) error
	ExistFile(context.Context, *File) (bool, error)
	FindFile(context.Context, *File) (*File, error)
}

// ToolsUsecase is a Tools usecase.
type ToolsUsecase struct {
	repo ToolsRepo
	log  *log.Helper
}

const filePrefix = "_file_"

// NewToolsUsecase new a Tools usecase.
func NewToolsUsecase(repo ToolsRepo, logger log.Logger) *ToolsUsecase {
	return &ToolsUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ToolsUsecase) UploadFile(ctx context.Context, req *v1.UploadFileRequest) (*v1.UploadFileReply, error) {
	//todo
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	var r strings.Builder
	_, err = r.Write(req.Content)
	if err != nil {
		return nil, err
	}
	_, err = r.WriteString(filePrefix)
	if err != nil {
		return nil, err
	}
	_, err = r.WriteString(fmt.Sprintf("%v", u.Id))
	if err != nil {
		return nil, err
	}
	uuid := common.MD5(r.String())
	exist, err := uc.repo.ExistFile(ctx, &File{ID: uuid})
	if err != nil {
		return nil, err
	}
	resp := &v1.UploadFileReply{Uuid: uuid}
	if exist {
		return resp, nil
	}
	//TODO 真实上传
	return resp, err
}
func (uc *ToolsUsecase) UploadFileByStream(ctx context.Context, req *v1.StreamRequest) (*v1.UploadFileReply, error) {
	//TODO 待调整
	u, err := constx.DefaultUser.User(ctx)
	if err != nil {
		return nil, err
	}
	var r strings.Builder
	_, err = r.Write(req.Chunk)
	if err != nil {
		return nil, err
	}
	_, err = r.WriteString(filePrefix)
	if err != nil {
		return nil, err
	}
	_, err = r.WriteString(fmt.Sprintf("%v", u.Id))
	if err != nil {
		return nil, err
	}
	uuid := common.MD5(r.String())
	exist, err := uc.repo.ExistFile(ctx, &File{ID: uuid})
	if err != nil {
		return nil, err
	}
	resp := &v1.UploadFileReply{Uuid: uuid}
	if exist {
		return resp, nil
	}
	//TODO 真实上传
	return resp, err
}
func (uc *ToolsUsecase) Files(ctx context.Context, req *global.IDStr) (*global.Byte, error) {
	//TODO 实现数据获取与读取逻辑
	return nil, nil
}
