package biz

import (
	"blog/api/global"
	v1 "blog/api/tools/v1"
	"blog/internal/common"
	"blog/internal/constx"
	"blog/pkg/aliyun/oss"
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/h2non/filetype"
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
	oss  *global.Oss
	log  *log.Helper
}

const filePrefix = "_file_"

// NewToolsUsecase new a Tools usecase.
func NewToolsUsecase(repo ToolsRepo, logger log.Logger, oss *global.Oss) *ToolsUsecase {
	return &ToolsUsecase{repo: repo, log: log.NewHelper(logger), oss: oss}
}

// UploadFile
// 文件在阿里云上存储的路径是 文件类型 /文件后缀 /年月日(主要是低频)/文件名称的形式
// 文件在本地的唯一ID是  文件内容的md5   //本来想是后者   文件内容 +固定后缀 _file_ +用户ID 的MD5  保证每个人有一份内容（主要是感觉这种场景比较少 没多大必要)
const filePath = "%v/%v/%v/%v"
const fileAuto = "auto"

func (uc *ToolsUsecase) UploadFile(ctx context.Context, req *v1.UploadFileRequest) (*v1.UploadFileReply, error) {
	var r strings.Builder
	_, err := r.Write(req.Content)
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
	kind, _ := filetype.Match(req.Content)
	ty := strings.Split(kind.MIME.Value, "/")[0]
	if kind == filetype.Unknown {
		extArr := strings.Split(req.Filename, ".")
		ext := extArr[len(extArr)-1]
		if ext == "" {
			ext = fileAuto
		} else {
			ok, _ := regexp.MatchString("^([A-z]+)$", ext)
			//如果不是纯英文 或者长度大于15 就单独命名了
			if !ok || len(ext) > 15 {
				ext = fileAuto
			}
		}
		kind = filetype.NewType(ext, fileAuto)
		ty = strings.Split(kind.MIME.Value, "/")[0]
	}
	path := fmt.Sprintf(filePath, ty, kind.Extension,
		time.Now().Format("20060102"), req.Filename)
	fileCli := oss.NewFileClient(ctx).SetBucket(uc.oss.Bucket).SetSite(uc.oss.Site).
		SetOSSClient(uc.oss.Key, uc.oss.Secret, uc.oss.Region)
	res, err := fileCli.UploadFile(path, req.Content)
	if err != nil {
		return nil, err
	}
	err = uc.repo.SaveFile(ctx, &File{
		ID:   uuid,
		Type: ty,
		Size: int64(len(req.Content)),
		Name: req.Filename,
		Path: res,
	})
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
