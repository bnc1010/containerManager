package auth

import (
	"context"
	"io/ioutil"
    "encoding/json"
	"path/filepath"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/bnc1010/containerManager/biz/utils"
)


type Routers struct {
    Root 			[]string  `json:"root"`
    Admin 			[]string  `json:"admin"`
    Common    		[]string  `json:"common"`
}

var AuthRouters *Routers


func InitAuthRouters() bool {
	ctx := context.Background()
	confPath := utils.GetConfAbPath()
	data, err := ioutil.ReadFile(filepath.Join(confPath, "routers.json"))
	if err != nil {
		hlog.CtxErrorf(ctx, "[AuthRouters] ReadInConfig failed, err: %v", err)
		return false
	}
	err = json.Unmarshal(data, &AuthRouters)
	if err != nil {
		hlog.CtxErrorf(ctx, "[AuthRouters] ReadInConfig failed, err: %v", err)
		return false
	}
	hlog.CtxInfof(ctx, "[AuthRouters] init success")
	// fmt.Println(AuthRouters)
	return true
}