package viper

import (
	"context"
	"os"
	"path/filepath"
	"github.com/bnc1010/containerManager/biz/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/spf13/viper"
)

var Conf *Config


func InitViper() bool {
	ctx := context.Background()
	viper.SetConfigType("yaml")
	runEnv := os.Getenv("RUN_ENV")
	confPath := utils.GetConfAbPath()
	if runEnv == "DEV" {
		viper.SetConfigFile(filepath.Join(confPath, "dev.config.yaml"))
	} else if runEnv == "PROD" {
		viper.SetConfigFile(filepath.Join(confPath, "prod.config.yaml"))
	} else {
		viper.SetConfigFile(filepath.Join(confPath, "default.config.yaml"))
	}

	if err := viper.ReadInConfig(); err != nil {
		hlog.CtxErrorf(ctx, "[Viper] ReadInConfig failed, err: %v", err)
		return false
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		hlog.CtxErrorf(ctx, "[Viper] Unmarshal failed, err: %v", err)
		return false
	}

	hlog.CtxInfof(ctx, "[Viper] Conf.App: %#v", Conf.App)
	hlog.CtxInfof(ctx, "[Viper] Conf.Cronjob: %#v", Conf.Cronjob)
	hlog.CtxInfof(ctx, "[Viper] Conf.Redis: %#v", Conf.Redis)
	hlog.CtxInfof(ctx, "[Viper] Conf.Postgres: %#v", Conf.Postgres)
	return true
}