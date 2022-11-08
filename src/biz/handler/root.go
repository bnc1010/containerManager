package handler

import (
	"fmt"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/bnc1010/containerManager/biz/pkg/k8s"
	resp_utils "github.com/bnc1010/containerManager/biz/utils"
)



func RootPostTest(ctx context.Context, c *app.RequestContext) {
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

func RootGetNamespaceList(ctx context.Context, c *app.RequestContext) {
	namespaceList, err := k8s.GetNamespaceList()
	if err != nil {
		resp_utils.ResponseError(c, "Get Namespaces Error", err)
		return
	}
	resp_utils.ResponseOK(c, "success", namespaceList)
}

func RootGetNamespace(ctx context.Context, c *app.RequestContext) {
	var req Namespace
    err := c.BindAndValidate(&req)
	
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	namespaceInfo, err := k8s.GetNamespace(req.Name)
	if err != nil {
		resp_utils.ResponseError(c, "Get Namespace Error", err)
		return
	}
	resp_utils.ResponseOK(c, "success", namespaceInfo)
}

func RootCreateNamespace(ctx context.Context, c *app.RequestContext) {
	var req Namespace
    err := c.BindAndValidate(&req)
	
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	namespaceInfo, err := k8s.CreateNamespace(req.Name)
	if err != nil {
		resp_utils.ResponseError(c, "Create Namespace Error", err)
		return
	}
	resp_utils.ResponseOK(c, "success", namespaceInfo)
}


func RootDeleteNamespace(ctx context.Context, c *app.RequestContext) {
	var req Namespace
    err := c.BindAndValidate(&req)
	
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	err = k8s.DeleteNamespace(req.Name)
	if err != nil {
		resp_utils.ResponseError(c, "Delete Namespace Error", err)
		return
	}
	resp_utils.ResponseOK(c, "success", "")
}



func RootGetNodeList(ctx context.Context, c *app.RequestContext) {
	nodeList, err := k8s.GetNodeList()
	if err != nil {
		resp_utils.ResponseError(c, "Get NodeList Error", err)
		return
	}
	resp_utils.ResponseOK(c, "success", nodeList)
}


func RootGetNode(ctx context.Context, c *app.RequestContext) {
	var req Node 
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	nodeInfo, err := k8s.GetNode(req.Name)
	if err != nil {
		resp_utils.ResponseError(c, "Get Node Error", err)
		return
	}
	resp_utils.ResponseOK(c, "success", nodeInfo)
}

func RootNodesMetrics(ctx context.Context, c *app.RequestContext) {
	data, err := k8s.NodesMetrics()
	if err != nil {
		resp_utils.ResponseError(c, "Get Nodes Metrics", err)
		return
	}
	resp_utils.ResponseOK(c, "success", data)
}