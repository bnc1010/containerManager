package postgres

import (
	"context"
	"fmt"
	viper "github.com/bnc1010/containerManager/biz/pkg/viper"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"database/sql"
	_ "github.com/lib/pq"
)

var (
	Client *sql.DB
	config *viper.Postgres
)


func InitPostgres() bool {
	config = viper.Conf.Postgres
	return initPostgres(context.Background(), &Client)
}

func initPostgres(ctx context.Context, client **sql.DB) bool {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",config.Host, config.Port, config.User, config.Password, config.Db)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil || db == nil{
		hlog.CtxFatalf(ctx, "[Postgres] Init Failed")
		return false
	}
	*client = db
	err = Client.Ping()
	if err != nil {
		hlog.CtxFatalf(ctx, "[Postgres] connect error")
		return false
	}
	Client.SetMaxOpenConns(20)
	Client.SetMaxIdleConns(10)
	hlog.CtxInfof(ctx, "[Postgres] Successfully connected!")
	return true
}