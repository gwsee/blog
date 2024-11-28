package data

import (
	"blog/app/account/internal/conf"
	"blog/internal/ent"
	_ "blog/internal/ent/runtime"
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAccountRepo)

// Data .
type Data struct {
	db       *ent.Client
	redisCli redis.Cmdable
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	dbCli, err := NewEntClient(c, logger)
	if err != nil {
		return nil, nil, err
	}
	redisCli, err := NewRedisCmd(c, logger)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		_ = dbCli.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: dbCli, redisCli: redisCli}, cleanup, nil
}
func NewEntClient(conf *conf.Data, logger log.Logger) (*ent.Client, error) {
	logs := log.NewHelper(log.With(logger, "module", "account/data/ent"))
	db, err := sql.Open(conf.Database.Driver, conf.Database.Source)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Duration(7200) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(600) * time.Second)
	var driver *entsql.Driver
	switch conf.Database.Driver {
	case "mysql":
		driver = entsql.OpenDB(dialect.MySQL, db)
	case "pgx/v5", "pgx":
		driver = entsql.OpenDB(dialect.Postgres, db)
	default:
		return nil, fmt.Errorf("unsupported database dialect: %s", conf.Database.Driver)
	}
	client := ent.NewClient(ent.Driver(driver), ent.Log(func(args ...any) {
		logs.Info("closing the data resources", args)
	}))
	return client, nil
}
func NewRedisCmd(conf *conf.Data, logger log.Logger) (redis.Cmdable, error) {
	logs := log.NewHelper(log.With(logger, "module", "account/data/redis"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Username:     conf.Redis.UserName,
		Password:     conf.Redis.Pws,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		logs.Errorf("redis connect error: %v", err)
		return client, nil
	}
	return client, err
}
