package data

import (
	"blog/internal/constx"
	gen "blog/internal/ent"
	"blog/internal/ent/files"
	"context"
	"time"

	"blog/app/tools/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type toolsRepo struct {
	data *Data
	log  *log.Helper
}

// NewToolsRepo .
func NewToolsRepo(data *Data, logger log.Logger) biz.ToolsRepo {
	return &toolsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *toolsRepo) SaveFile(ctx context.Context, g *biz.File) error {
	return r.data.db.Files.Create().
		SetID(g.ID).SetName(g.Name).SetType(g.Type).SetSize(g.Size).SetPath(g.Path).
		OnConflict().DoNothing().Exec(ctx)
}
func (r *toolsRepo) ExistFile(ctx context.Context, g *biz.File) (string, error) {
	f, err := r.data.db.Files.Query().Where(files.IDEQ(g.ID)).Only(ctx)
	if err != nil {
		if gen.IsNotFound(err) {
			return "", nil
		}
		return "", err
	}
	return f.Path, nil
}
func (r *toolsRepo) FindFile(ctx context.Context, g *biz.File) (*biz.File, error) {
	file, err := r.data.db.Files.Get(ctx, g.ID)
	if err != nil {
		return nil, err
	}
	return &biz.File{
		ID:   file.ID,
		Type: file.Type,
		Size: file.Size,
		Name: file.Name,
		Path: file.Path,
	}, nil
}
func (r *toolsRepo) CacheFile(ctx context.Context, origin, path string, expire int64) {
	if r.data.redisCli == nil {
		return
	}
	r.data.redisCli.Set(ctx, constx.FileCachePrefix+origin, path, time.Duration(expire)*time.Minute)
}
