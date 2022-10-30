package viper

type Config struct {
	App     *App     `yaml:"App"`
	Cronjob *Cronjob `yaml:"Cronjob"`
}

type App struct {
	HostPorts          string `yaml:"HostPorts"`          // 服务监听的地址和端口
	MaxRequestBodySize int    `yaml:"MaxRequestBodySize"` // 最大的请求体大小
}

type Cronjob struct {
	TempFileMinute float64 `yaml:"TempFileMinute"` // 文件上传token刷新时间（默认1h过期）
	TokenMinute    float64 `yaml:"TokenMinute"`    // 临时文件夹最长生存时间
}