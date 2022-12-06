package test

import (
	"fmt"
	"github.com/bnc1010/containerManager/biz/utils"
)


func TestJson(){
	var mp_a = map[string] interface{} {"limits": map[string] interface{} {"cpu":"500m", "memory":"1000Mi"}}
	fmt.Println(mp_a)
	ss, err := utils.Map2Bytes(mp_a)
	if err != nil {
		fmt.Println(err)
	}
	mp_b, err := utils.Bytes2Map(ss)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mp_b)
	}
}