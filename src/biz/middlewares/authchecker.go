package middlewares

import (
	"os"
	"fmt"
	"errors"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/bnc1010/containerManager/biz/utils"
	viper "github.com/bnc1010/containerManager/biz/pkg/viper"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	
)

var (
	secretKey []byte
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
	secretKey = []byte(strKey)
	hlog.CtxInfof(context.Background(), "[Checker] init success")
	return true
}


//curl -d '{"a":"test"}' -H "Content-Type:application/json" -H "AUTH_TOKEN:testtoken"  -X POST http://127.0.0.1:8888/postping.1:8888/postping
//暂时使用简单对称加密
func checkToken(c *app.RequestContext) {
	AUTH_TOKEN := string(c.GetHeader(viper.Conf.App.TokenHeader))
	fmt.Println(AUTH_TOKEN)
	encrypted  := utils.AesEncryptCBC([]byte("hello"), secretKey)
	bs64 := utils.Base64Encoding(encrypted)
	fmt.Println(string(bs64))
	de64 := utils.Base64Decoding(bs64)
	content := utils.AesDecryptCBC([]byte(de64), secretKey)
	fmt.Println(string(content))
}

//进入handler之前预鉴定
func AuthChecker() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
	  // pre-handle
	  fmt.Println("enter authchecker") 
	  checkToken(c)
	  fmt.Println(c.FullPath())
	  fmt.Println(string(c.URI().Path()))
	  
	  c.Next(ctx)
	  // post-handle
	}
  }