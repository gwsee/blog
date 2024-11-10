package data

import (
	"blog/app/account/internal/conf"
	"blog/internal/ent"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	_ "blog/internal/ent/runtime"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAccountRepo)

// Data .
type Data struct {
	//// TODO wrapped database client
	//*entsql.Driver
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		return nil, nil, err
	}
	//todo 下面的动态设置进入配置
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Duration(7200) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(600) * time.Second)
	cleanup := func() {
		_ = db.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	var driver *entsql.Driver
	switch c.Database.Driver {
	case "mysql":
		driver = entsql.OpenDB(dialect.MySQL, db)
	case "pgx/v5", "pgx":
		driver = entsql.OpenDB(dialect.Postgres, db)
	default:
		return nil, nil, fmt.Errorf("unsupported database dialect: %s", c.Database.Driver)
	}
	client := ent.NewClient(ent.Driver(driver), ent.Log(func(args ...any) {
		//logz.Info(ctx, "ent log", logz.Any("args", args))
		log.NewHelper(logger).Info("closing the data resources", args)
	}))
	return &Data{client}, cleanup, nil
	//switch c.Database.Driver {
	//case "mysql":
	//	return &Data{entsql.OpenDB(dialect.MySQL, db)}, cleanup, nil
	//case "pgx/v5", "pgx":
	//	return &Data{entsql.OpenDB(dialect.Postgres, db)}, cleanup, nil
	//default:
	//	return nil, nil, fmt.Errorf("unsupported database dialect: %s", c.Database.Driver)
	//}
	//return &Data{}, cleanup, nil
}
