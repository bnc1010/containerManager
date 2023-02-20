package middlewares

import (
	"os"
	"fmt"
	"errors"
	"regexp"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/bnc1010/containerManager/biz/utils"
	viper "github.com/bnc1010/containerManager/biz/pkg/viper"
	auth "github.com/bnc1010/containerManager/biz/pkg/auth"
	customError "github.com/bnc1010/containerManager/biz/pkg/error"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	
)

var (
	SecretKey []byte
)
// c.FullPath()
// c.GetHeader()


func keyAppend(k string) (string, error) {
	lenK := len(k)
	if lenK <= 32 {
		for i:=lenK; i < 32 ; i++ {
			k = k + "*"
		}
		return k, nil
	}else{
		errorInfo := fmt.Sprintf("%s", "[Checker] SecretKey too long, max length is 32")
		hlog.CtxErrorf(context.Background(), errorInfo)
		return "",  errors.New(errorInfo)
	}
}

func InitChecker() bool {
	strKey, err := keyAppend(os.Getenv("SECRET_KEY"))
	if err != nil {
		return false
	}
	SecretKey = []byte(strKey)
	hlog.CtxInfof(context.Background(), "[Checker] init success")
	return true
}

func DecodeToken(AUTH_TOKEN string) (content []byte) {
	defer func() {
		err := recover()
		if err != nil {
			hlog.CtxErrorf(context.Background(), "token decode error")
		}
	}()
	de64 := utils.Base64Decoding(AUTH_TOKEN)
	res := utils.AesDecryptCBC([]byte(de64), SecretKey)
	return res
}


//curl -d '{"a":"test"}' -H "Content-Type:application/json" -H "AUTH_TOKEN:grFsvAdxNlb6YdY1e5nz1o3gQ89tmrFzVAotNW00ZD8="  -X POST http://127.0.0.1:8888/postping
//暂时使用简单对称加密
func checkToken(c *app.RequestContext) (token * auth.Token, errorMsg string) {
	AUTH_TOKEN := string(c.GetHeader(viper.Conf.App.TokenHeader))
	content := DecodeToken(AUTH_TOKEN)
	if content == nil {
		return nil, fmt.Sprintf("Illegal token!")
	}
	token, error := auth.GenerateFromStr(string(content))
	switch error.(type) {
		case nil:
			return token, ""
		case * customError.TokenTimeoutError:
			return nil, fmt.Sprintf("%s", error)
		default:
			return nil, fmt.Sprintf("Illegal token!")
	}
	return nil, fmt.Sprintf("Illegal token!")
}

func checkRouter(token * auth.Token, c *app.RequestContext) (bool, string) {
	var routers []string
		switch token.Role {
			case "root":
				routers = auth.AuthRouters.Root
			case "admin":
				routers = auth.AuthRouters.Admin
			case "common":
				routers = auth.AuthRouters.Common
			default:
				routers = nil
		}
		if routers == nil {
			return false, "Error Role!"
		}

		for _, router := range routers {
			sta, err := regexp.MatchString(router, c.FullPath())
			if err == nil && sta {
				return true, ""
			}
		}
		return false, "No Permission!"
}

//进入handler之前预鉴定
func AuthChecker() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		//
		// 检查地址是否存在
		//
		if c.FullPath() == "" {
			utils.ResponseNotFound(c)
			c.AbortWithStatus(404)
			return
		}
		
		//
		// 检查token是否有效:能否解密&是否过期
		//
		token, errorMsg := checkToken(c)
		if token == nil {
			utils.ResponseForbid(c, errorMsg)
			c.AbortWithStatus(403)
			return
		}

		//
		// 检查token是否对访问的地址拥有权限
		//
		sta, errorMsg := checkRouter(token, c)
		if !sta {
			utils.ResponseForbid(c, errorMsg)
			c.AbortWithStatus(403)
			return
		}

		// 传递当前访问用户的id给业务handler   使用ctx.Value("requestUserId")获取该值
		ctx = context.WithValue(ctx, "requestUserId", token.UserId)
		ctx = context.WithValue(ctx, "requestUserRole", token.Role)
		c.Next(ctx)
	}
  }