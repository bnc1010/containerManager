package main

import (
	"os"
    viper "github.com/bnc1010/containerManager/biz/pkg/viper"
    redis "github.com/bnc1010/containerManager/biz/pkg/redis"
    godotenv "github.com/bnc1010/containerManager/biz/pkg/godotenv"
    k8s "github.com/bnc1010/containerManager/biz/pkg/k8s"
    "github.com/cloudwego/hertz/pkg/app/server"
    test "github.com/bnc1010/containerManager/biz/test"
    middlewares "github.com/bnc1010/containerManager/biz/middlewares"
)

func main() {
	Init()
    // Test()
    
    ServerStart()
}

func Test() {
    k8s.Test()
    test.EncryptionTest()
    test.TestToken()
}

func Init() {
    initSignal := true
    initSignal = initSignal && godotenv.InitGodotenv()
    initSignal = initSignal && viper.InitViper()
    if os.Getenv("USE_REDIS") == "TRUE" {
        initSignal = initSignal && redis.InitRedis()
    }
    initSignal = initSignal && k8s.InitK8s()
    initSignal = initSignal && middlewares.InitChecker()
    if ! initSignal {
        os.Exit(101)
    }
}

func ServerStart() {
    h := server.Default(
        server.WithHostPorts(viper.Conf.App.HostPorts),
		server.WithMaxRequestBodySize(viper.Conf.App.MaxRequestBodySize),
		server.WithExitWaitTime(10),
    )
    h.Use(middlewares.IpChecker())
    h.Use(middlewares.AuthChecker())
	register(h)
	h.Spin()
}