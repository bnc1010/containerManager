package test

import (
	"fmt"
	// "time"
	"github.com/bnc1010/containerManager/biz/pkg/auth"
	"github.com/bnc1010/containerManager/biz/utils"
	"github.com/bnc1010/containerManager/biz/middlewares"
)



func TestToken() {
	token, err := auth.GenerateFromStr("bnc&423h4huhuhfuseu34&1667213184&admin")  
	if err == nil {
		fmt.Println(token)
		fmt.Println(token.UserName)
		fmt.Println(token.UserId)
		fmt.Println(token.CreateTime)
		fmt.Println(token.Role)
	}else {
		fmt.Println("illegal token", err)
	}

	encrypted  := utils.AesEncryptCBC([]byte("bnc&423h4huhuhfuseu34&1667213184&admin"), middlewares.SecretKey)
	bs64 := utils.Base64Encoding(encrypted)
	fmt.Println(bs64)
}

//curl -d '{"a":"test"}' -H "Content-Type:application/json" -H "AUTH_TOKEN:grFsvAdxNlb6YdY1e5nz1o3gQ89tmrFzVAotNW00ZD8="  -X POST http://127.0.0.1:8888/postping
//Aa2N9jIOFz4If8Qn/EPGAn2nTd4z0BkcM45E6YetcGI1x9NOgDkUQFftPcNaAI6R   root
//Aa2N9jIOFz4If8Qn/EPGAn2nTd4z0BkcM45E6YetcGKcS7TGwPycN6/JRWUFjSEZ   common
//Aa2N9jIOFz4If8Qn/EPGAn2nTd4z0BkcM45E6YetcGJojft0Yh9qG2UT7RdUNvGY	 admin