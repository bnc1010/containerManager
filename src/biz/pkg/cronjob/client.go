package cronjob

import (
	"context"
	"math/rand"
	"time"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/pkg/viper"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/xhdd123321/whicinth-steganography-bd/biz/service/fileService"
)

var config *viper.Cronjob

// InitCronjob 初始化定时任务
func InitCronjob() {
	// 配置初始化
	config = viper.Conf.Cronjob

	ctx := context.Background()
	go expiredFileClearCronJob(ctx)
	hlog.CtxInfof(ctx, "[Cronjob] ExpiredFileClearCronJob start...")
}

// expiredFileClearCronJob 过期文件清理任务
func expiredFileClearCronJob(ctx context.Context) {
	// 启动时先执行一遍清理
	err := fileService.ClearFile(ctx, config.TempFileMinute)
	if err != nil {
		hlog.CtxErrorf(ctx, "[Cronjob] ClearFile cronjob run failed, err: %v", err)
	}
	// 启动定时清理任务
	sched := time.Tick(time.Minute * time.Duration(config.TempFileMinute))
	for range sched {
		err := fileService.ClearFile(ctx, config.TempFileMinute)
		if err != nil {
			hlog.CtxErrorf(ctx, "[Cronjob] ClearFile cronjob run failed, err: %v", err)
		}
	}
}
