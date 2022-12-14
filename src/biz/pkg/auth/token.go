package auth

import (
	"os"
	"time"
	"errors"
	"context"
	"strings"
	"strconv"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	customError "github.com/bnc1010/containerManager/biz/pkg/error"
)

const (
	tokenTimeout       time.Duration = 30 * time.Minute
)

type Token struct {
	UserName        	string
	UserId				string
	CreateTime			time.Time
	Role				string
}


//判断token是否过期
func (token *Token) IsValid() bool {
	existTime := time.Since(token.CreateTime)
	return existTime <= tokenTimeout
}
//
// TokenStr :   bob&1667213184&root
//				username&timestamp&role
//
func GenerateFromStr(tokenStr string) (*Token, error) {
	defer func() {
		err := recover()
		if err != nil {
			hlog.CtxErrorf(context.Background(), "parameters error")
		}
	}()

	infos := strings.Split(tokenStr, "&")
	if len(infos) != 4 {
		return nil, errors.New("parameters error")
	}
	
	timeUnix, err := strconv.ParseInt(infos[2],10,64)
	if err != nil{
		return nil, err
	}

	tm := time.Unix(timeUnix, 0)
	
	maxTime, err := strconv.ParseInt(os.Getenv("TOKEN_TIMEOUT"),10,64)
	if err != nil {
		maxTime = 3600
	}
	if maxTime != -1 {
		if time.Now().Unix() - tm.Unix() >= 3600 {
			return nil, &customError.TokenTimeoutError{infos[0], tm, infos[3]}
		}
	}
	
	token := &Token{infos[0], infos[1], tm, infos[3]}
	return token, nil
}