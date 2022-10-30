package main

import (
	viper "github.com/bnc1010/containerManager/biz/pkg/viper"
    godotenv "github.com/bnc1010/containerManager/biz/pkg/godotenv"
    k8s "github.com/bnc1010/containerManager/biz/pkg/k8s"
    "github.com/cloudwego/hertz/pkg/app/server"
    test "github.com/bnc1010/containerManager/biz/test"
)

func main() {
	Init()
    // Test()
    
    ServerStart()
}

func Test() {
    k8s.Test()
    test.EncryptionTest()
}

func Init() {
    godotenv.InitGodotenv()
    viper.InitViper()
    k8s.InitK8s()
}

func ServerStart() {
    h := server.Default()
	register(h)
	h.Spin()
}