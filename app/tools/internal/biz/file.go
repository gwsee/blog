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
	ExistFile(context.Context, *File) (string, error)
	FindFile(context.Context, *File) (*File, error)
	CacheFile(context.Context, string, string, int64)
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

func (uc *ToolsUsecase) UploadFile(ctx context.Context, req *v1.UploadFileRequest) (resp *v1.UploadFileReply, err error) {
	var r strings.Builder
	_, err = r.Write(req.Content)
	if err != nil {
		return nil, err
	}
	uuid := common.MD5(r.String())
	var exist string

	//返回的时候将临时链接返回
	defer func() {
		if err != nil {
			return
		}
		c, _ := uc.fileLink(ctx, uuid, exist)
		if c == nil {
			return
		}
		resp.Message = c.Id
	}()

	//判断文件是否存在
	exist, err = uc.repo.ExistFile(ctx, &File{ID: uuid})
	if err != nil {
		return nil, err
	}
	resp = &v1.UploadFileReply{Uuid: uuid}
	if exist != "" {
		return resp, nil
	}
	//自定义存储路径
	kind, _ := filetype.Match(req.Content)
	ty := strings.Split(kind.MIME.Value, "/")[0]
	extArr := strings.Split(req.Filename, ".")
	if kind == filetype.Unknown {
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
	//md5文件名称到oss
	extArr[0] = uuid
	path := fmt.Sprintf(filePath, ty, kind.Extension,
		time.Now().Format("20060102"), strings.Join(extArr, "."))

	//上传文件并记录关系
	fileCli := oss.NewFileClient(ctx).SetBucket(uc.oss.Bucket).SetSite(uc.oss.Site).
		SetOSSClient(uc.oss.Key, uc.oss.Secret, uc.oss.Region)
	exist, err = fileCli.UploadFile(path, req.Content)
	if err != nil {
		return nil, err
	}
	err = uc.repo.SaveFile(ctx, &File{
		ID:   uuid,
		Type: ty,
		Size: int64(len(req.Content)),
		Name: req.Filename,
		Path: exist,
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
	if exist != "" {
		resp.Message = exist
		return resp, nil
	}
	//TODO 真实上传
	return resp, err
}
func (uc *ToolsUsecase) Files(ctx context.Context, req *global.IDStr) (*global.IDStr, error) {
	one, err := uc.repo.FindFile(ctx, &File{ID: req.Id})
	if err != nil {
		return nil, err
	}
	return uc.fileLink(ctx, req.Id, one.Path)
}
func (uc *ToolsUsecase) UploadOssToken(ctx context.Context, req *global.IDStr) (*v1.UploadOssTokenReply, error) {
	//判断文件是否存在
	exist, err := uc.repo.ExistFile(ctx, &File{ID: req.Id})
	if err != nil {
		return nil, err
	}
	if exist != "" {
		return &v1.UploadOssTokenReply{Exist: true}, nil
	}
	//不存在就获取token
	res, err := oss.NewFileClient(ctx).UploadToken(uc.oss)
	if err != nil {
		return nil, err
	}
	return &v1.UploadOssTokenReply{
		Host:      uc.oss.Site,
		Expire:    uc.oss.Expire,
		Region:    uc.oss.Region,
		Bucket:    uc.oss.Bucket,
		ExpireStr: *res.Expiration,
		KeyId:     *res.AccessKeyId,
		KeySecret: *res.AccessKeySecret,
		Token:     *res.SecurityToken,
	}, nil
}
func (uc *ToolsUsecase) UploadOssSave(ctx context.Context, req *v1.UploadOssSaveReq) (*global.Empty, error) {
	return &global.Empty{}, uc.repo.SaveFile(ctx, &File{
		ID:   req.Id,
		Type: req.Type,
		Size: req.Size,
		Name: req.Name,
		Path: req.Path,
	})
}
func (uc *ToolsUsecase) fileLink(ctx context.Context, id, url string) (*global.IDStr, error) {
	fileCli := oss.NewFileClient(ctx).SetBucket(uc.oss.Bucket).SetSite(uc.oss.Site).
		SetOSSClient(uc.oss.Key, uc.oss.Secret, uc.oss.Region)
	link, _ := fileCli.Temp(url, int(uc.oss.Expire))
	expire := uc.oss.Expire
	if link == "" {
		expire = 10
	}
	uc.repo.CacheFile(ctx, id, link, expire)
	return &global.IDStr{Id: link}, nil
}
