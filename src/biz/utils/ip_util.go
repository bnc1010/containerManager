package utils

import (
	"net"
	"time"
	"strings"
	"strconv"
	"math/rand"
	"github.com/cloudwego/hertz/pkg/app"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

// RemoteIp 返回远程客户端的 IP，如 192.168.1.1
func RemoteIp(c *app.RequestContext) string {
	remoteAddr := c.RemoteAddr().String()
	if ip := string(c.GetHeader(XRealIP)); ip != "" {
		remoteAddr = ip
	} else if ip = string(c.GetHeader(XForwardedFor)); ip != "" {
		remoteAddr = ip
	}

	if remoteAddr == "" {
		remoteAddr = "127.0.0.1"
	}
	remoteAddr = strings.Split(remoteAddr, ":")[0]
	return remoteAddr
}


func ScanPort(protocol string, hostname string, port int) bool {
	p := strconv.Itoa(port)
	addr := net.JoinHostPort(hostname, p)
	conn, err := net.DialTimeout(protocol, addr, 200*time.Millisecond)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func RandUsablePort(hostname string) int32 {
    rand.Seed(time.Now().UnixNano())
	tryTime := 0
    for ; ; {
		port := rand.Intn(1000) + 30000
		if !ScanPort("tcp", hostname, port){
			return int32(port)
		}
		tryTime = tryTime + 1
		if tryTime == 10 {
			return -1
		}
	}
}