package test

import (
	"fmt"
	"github.com/bnc1010/containerManager/biz/utils"
)


func TestPort(){
	fmt.Println(utils.ScanPort("tcp", "47.100.69.138", 30257))
	// fmt.Println(utils.ScanPort("udp", "47.100.69.138", 30257))
	fmt.Println(utils.RandUsablePort("47.100.69.138"))
}