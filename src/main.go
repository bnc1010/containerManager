package main

import (
	"os"
    viper "github.com/bnc1010/containerManager/biz/pkg/viper"
    redis "github.com/bnc1010/containerManager/biz/pkg/redis"
    godotenv "github.com/bnc1010/containerManager/biz/pkg/godotenv"
    k8s "github.com/bnc1010/containerManager/biz/pkg/k8s"
    filecontrol "github.com/bnc1010/containerManager/biz/pkg/filecontrol"
    "github.com/cloudwego/hertz/pkg/app/server"
    test "github.com/bnc1010/containerManager/biz/test"
    middlewares "github.com/bnc1010/containerManager/biz/middlewares"
    auth "github.com/bnc1010/containerManager/biz/pkg/auth"
    postgres "github.com/bnc1010/containerManager/biz/pkg/postgres"
)

func main() {
	Init()
    Test()
    
    // ServerStart()
}

//
// 自定义测试项
//
func Test() {
    // k8s.Test()
    // test.EncryptionTest()
    // test.TestToken()
    // test.TestK8s()
    test.TestPostgres()
    // test.TestJson()
}

//
// 配置初始化
//
func Init() {
    initSignal := true
    initSignal = initSignal && godotenv.InitGodotenv()
    initSignal = initSignal && viper.InitViper()
    initSignal = initSignal && filecontrol.InitPath()
    if os.Getenv("USE_REDIS") == "TRUE" {
        initSignal = initSignal && redis.InitRedis()
    }
    initSignal = initSignal && k8s.InitK8s()
    initSignal = initSignal && middlewares.InitChecker()
    initSignal = initSignal && auth.InitAuthRouters()
    initSignal = initSignal && postgres.InitPostgres()
    if ! initSignal {
        os.Exit(101)
    }
}

//
// 启动服务
//
func ServerStart() {
    h := server.Default(
        server.WithHostPorts(viper.Conf.App.HostPorts),
		server.WithMaxRequestBodySize(viper.Conf.App.MaxRequestBodySize),
		server.WithExitWaitTime(10),
    )
    
//
// 添加中间件
//
    h.Use(middlewares.IpChecker())
    h.Use(middlewares.AuthChecker())

//
// 注册路由并启动服务
//
	register(h)
	h.Spin()
}