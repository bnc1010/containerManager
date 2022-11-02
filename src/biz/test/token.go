package test

import (
	"fmt"
	// "time"
	"github.com/bnc1010/containerManager/biz/pkg/auth"
	"github.com/bnc1010/containerManager/biz/utils"
	"github.com/bnc1010/containerManager/biz/middlewares"
)



func TestToken() {
	token, error := auth.GenerateFromStr("bnc&1667213184&common")  
	if error == nil {
		fmt.Println(token)
		fmt.Println(token.UserName)
		fmt.Println(token.CreateTime)
		fmt.Println(token.Role)
	}else {
		fmt.Println("illegal token")
	}

	encrypted  := utils.AesEncryptCBC([]byte("bnc&1667213184&common"), middlewares.SecretKey)
	bs64 := utils.Base64Encoding(encrypted)
	fmt.Println(bs64)
}

//curl -d '{"a":"test"}' -H "Content-Type:application/json" -H "AUTH_TOKEN:grFsvAdxNlb6YdY1e5nz1o3gQ89tmrFzVAotNW00ZD8="  -X POST http://127.0.0.1:8888/postping
//grFsvAdxNlb6YdY1e5nz1o3gQ89tmrFzVAotNW00ZD8=   root
//HkilzznwLrdwCbvwSo/+U4M1eJ1QXgQflLSkub5/lAY=   common