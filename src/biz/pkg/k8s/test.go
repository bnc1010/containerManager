package k8s

import (
	// apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"context"
	"fmt"
)


func testPod(ctx context.Context) {
	pods, err := Client.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		hlog.CtxErrorf(ctx, "[K8S] test error, err: %v", err)
	}
	for _,v := range  pods.Items {
        fmt.Printf("命名空间是：%v pod名字：%v\n",v.Namespace,v.Name)
    }
}

func testNode(ctx context.Context) {
	nodeList,err := GetNodeList()
	for err != nil{
		hlog.CtxErrorf(ctx, "[K8S] test error, err: %v", err)
	}
	//fmt.Printf("%+v",nodeList)
	for _,nodeInfo := range nodeList.Items{
		fmt.Printf("node 的名字为：|%s|\n",nodeInfo.Name)
	}

	nodeInfo, err := GetNode("cn-shanghai.10.24.0.129")
	if err == nil {
		fmt.Printf("%+v\n", nodeInfo.Status.Allocatable.Memory().String())
	} else {
		fmt.Printf("%s\n", err)
	}
}

func Test()  {
	ctx := context.Background()
	testPod(ctx)
	testNode(ctx)
}