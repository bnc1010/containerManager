package k8s

import (
	// apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"context"
	"fmt"
)

func Test()  {
	ctx := context.Background()
	pods, err := (*Client).CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		hlog.CtxErrorf(ctx, "[K8S] test error, err: %v", err)
	}
	for _,v := range  pods.Items {
        fmt.Println("命名空间是：%v\n pod名字：%v",v.Namespace,v.Name)
    }
}