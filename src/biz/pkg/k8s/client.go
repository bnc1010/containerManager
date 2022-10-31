package k8s

import (
	"path/filepath"
	"k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/kubernetes"
	"github.com/bnc1010/containerManager/biz/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"context"
)

var Client * kubernetes.Clientset

func InitK8s() bool{
	ctx := context.Background()
	confPath := utils.GetConfAbPath()
	kubeconfig := filepath.Join(confPath, "k8sconfig")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		hlog.CtxErrorf(ctx, "[K8S] config init failed, err: %v", err)
		return false
	}
	clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        hlog.CtxErrorf(ctx, "[K8S] client init failed, err: %v", err)
		return false
    }
	hlog.CtxInfof(ctx, "[K8S] init success")
	Client = clientset
	return true
}

func GetClient() * kubernetes.Clientset {
	return Client
}