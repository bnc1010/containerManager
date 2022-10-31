package godotenv

import (
	"context"
	"os"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/joho/godotenv"
)

// InitGodotenv 初始化环境变量
func InitGodotenv() bool {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		hlog.CtxErrorf(ctx, "[Godotenv] Init Env failed, err: %v", err)
		return false
	}
	hlog.CtxInfof(ctx, "[Godotenv] Init Env success, RUN_ENV: %v", os.Getenv("RUN_ENV"))
	hlog.CtxInfof(ctx, "[Godotenv] Init Env success, SECRET_KEY: %v", os.Getenv("SECRET_KEY"))
	return true
}