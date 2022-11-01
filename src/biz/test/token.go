package test

import (
	"fmt"
	// "time"
	"github.com/bnc1010/containerManager/biz/pkg/auth"
	"github.com/bnc1010/containerManager/biz/utils"
	"github.com/bnc1010/containerManager/biz/middlewares"
)



func TestToken() {
	// token := auth.Token{UserName:"bnc", CreateTime:time.Now()}
	// fmt.Println(token)
	// time.Sleep(time.Duration(2)*time.Second)
	// fmt.Println(token.IsValid())
	//1667213184
	// fmt.Println(time.Now().Unix())
	token, error := auth.GenerateFromStr("bnc&1667213184&root")  
	if error == nil {
		fmt.Println(token)
		fmt.Println(token.UserName)
		fmt.Println(token.CreateTime)
		fmt.Println(token.Role)
	}else {
		fmt.Println("illegal token")
	}

	encrypted  := utils.AesEncryptCBC([]byte("bnc&1667213184&root"), middlewares.SecretKey)
	bs64 := utils.Base64Encoding(encrypted)
	fmt.Println(bs64)
}

//grFsvAdxNlb6YdY1e5nz1o3gQ89tmrFzVAotNW00ZD8=