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
	resp_utils.ResponseOK(c, responseMsg.Success, podList)
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
	resp_utils.ResponseOK(c, responseMsg.Success, podInfo)
}

func AdminPodsMetrics(ctx context.Context, c *app.RequestContext) {
	data, err := k8s.PodsMetrics()
	if err != nil {
		resp_utils.ResponseError(c, "Get Pods Metrics", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, data)
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
	resp_utils.ResponseOK(c, responseMsg.Success, podUsage)
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
	resp_utils.ResponseOK(c, responseMsg.Success, podUsage)
}

func AdminPodLog(ctx context.Context, c *app.RequestContext) {
	var req PodLog
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	if !req.VaildTailLines() {
		resp_utils.ResponseErrorParameter(c, "lines should in [100,10000], and mod 100 is 0")
		return
	}
	logs, err := k8s.PodLog(req.Namespace, req.Pod, nil, &sinceTime, &req.TailLines)
	if err != nil {
		resp_utils.ResponseError(c, "Get Pod Log Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, logs)
}

// func AdminProjectList(ctx context.Context, c *app.RequestContext) {

// 	resp_utils.ResponseOK(c, responseMsg.Success)
// }

// func AdminGetProject(ctx context.Context, c *app.RequestContext) {
	
// 	resp_utils.ResponseOK(c, responseMsg.Success)
// }