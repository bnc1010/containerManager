package handler

import (
	"fmt"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/bnc1010/containerManager/biz/pkg/k8s"
	resp_utils "github.com/bnc1010/containerManager/biz/utils"
)

func AdminPostTest(ctx context.Context, c *app.RequestContext) {
	type Test struct {
		A string `json:"a" vd:"$!='Hertz'"`
	}
	var req Test
    err := c.BindAndValidate(&req)
	
	if err == nil {
		fmt.Println(req)
	} else{
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
	}
	
	c.JSON(200, utils.H{
		"message": "pong",
	})
}


func AdminGetPodList(ctx context.Context, c *app.RequestContext) {
	var req Namespace
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	podList,err := k8s.GetPodList(req.Name)
	if err != nil {
		resp_utils.ResponseError(c, "Get PodList Error", err)
		return
	}
	resp_utils.ResponseOK(c, "success", podList)
}


func AdminGetPod(ctx context.Context, c *app.RequestContext) {
	var req Pod
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	podInfo, err := k8s.GetPod(req.Namespace, req.Name)
	if err != nil {
		resp_utils.ResponseError(c, "Get Pod Error", err)
		return
	}
	resp_utils.ResponseOK(c, "success", podInfo)
}

func AdminPodsMetrics(ctx context.Context, c *app.RequestContext) {
	data, err := k8s.PodsMetrics()
	if err != nil {
		resp_utils.ResponseError(c, "Get Pods Metrics", err)
		return
	}
	resp_utils.ResponseOK(c, "success", data)
}

func AdminPodHeapsterMemory(ctx context.Context, c *app.RequestContext) {
	var req Pod
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	podUsage, err := k8s.PodHeapsterMemory(req.Namespace, req.Name)
	if err != nil {
		resp_utils.ResponseError(c, "Get Pod Memory Info Error", err)
		return
	}
	resp_utils.ResponseOK(c, "success", podUsage)
}

func AdminPodHeapsterCpu(ctx context.Context, c *app.RequestContext) {
	var req Pod
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	podUsage, err := k8s.PodHeapsterCpu(req.Namespace, req.Name)
	if err != nil {
		resp_utils.ResponseError(c, "Get Pod Cpu Info Error", err)
		return
	}
	resp_utils.ResponseOK(c, "success", podUsage)
}