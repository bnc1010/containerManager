package middlewares

import (
	"os"
	"fmt"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/bnc1010/containerManager/biz/utils"
	// "github.com/cloudwego/hertz/pkg/common/hlog"
	
)

func frequencyControl(ip string) {
	//控制单个ip的访问频率
	if os.Getenv("USE_REDIS") == "TRUE" && os.Getenv("WHITE_LIST") == "TRUE" {
		//todo: 白名单直接放行
	}
	return 	
}

func checkIp(c *app.RequestContext) {
	clientIp := utils.RemoteIp(c)
	
	if os.Getenv("USE_REDIS") == "TRUE" && os.Getenv("BLACK_LIST") == "TRUE" {
		//todo: 黑名单控制
		
	}
	
	fmt.Println(clientIp)

	if os.Getenv("USE_REDIS") == "TRUE" && os.Getenv("IP_FREQUENCY_CONTROL") == "TRUE" {
		frequencyControl(clientIp)
	}
}


func IpChecker() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// pre-handle
		fmt.Println("enter ipchecker")
		checkIp(c)
		
		c.Next(ctx)
		// post-handle
	  }
}