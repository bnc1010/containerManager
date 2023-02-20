package handler

import (
	"fmt"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/bnc1010/containerManager/biz/pkg/k8s"
	"github.com/bnc1010/containerManager/biz/pkg/postgres"
	resp_utils "github.com/bnc1010/containerManager/biz/utils"
)



func RootPermissionCheck(ctx context.Context, c *app.RequestContext) {
	requestUserId := fmt.Sprintf("%v",ctx.Value("requestUserId"))
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
	res_data := make(map[string]string)
	res_data["userId"] = requestUserId
	resp_utils.ResponseOK(c, "ok", res_data)
}

func RootGetNamespaceList(ctx context.Context, c *app.RequestContext) {
	namespaceList, err := k8s.GetNamespaceList()
	if err != nil {
		resp_utils.ResponseError(c, "Get Namespaces Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, namespaceList)
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
	resp_utils.ResponseOK(c, responseMsg.Success, namespaceInfo)
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
	resp_utils.ResponseOK(c, responseMsg.Success, namespaceInfo)
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
	resp_utils.ResponseOK(c, responseMsg.Success, "")
}



func RootGetNodeList(ctx context.Context, c *app.RequestContext) {
	nodeList, err := k8s.GetNodeList()
	if err != nil {
		resp_utils.ResponseError(c, "Get NodeList Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, nodeList)
}


func RootGetNode(ctx context.Context, c *app.RequestContext) {
	var req Node 
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	fmt.Println(req)
	nodeInfo, err := k8s.GetNode(req.Name)
	
	if err != nil {
		resp_utils.ResponseError(c, "Get Node Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, nodeInfo)
}

func RootNodesMetrics(ctx context.Context, c *app.RequestContext) {
	data, err := k8s.NodesMetrics()
	if err != nil {
		resp_utils.ResponseError(c, "Get Nodes Metrics", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, data)
}

func RootAddK8sNodeTag(ctx context.Context, c *app.RequestContext) {
	var req NodeTag
  err := c.BindAndValidate(&req)
	if err != nil {
		resp_utils.ResponseErrorParameter(c)
		return 
	}
	nodeTagId := resp_utils.RandStringWithLengthN(36)
	nodeTag := &postgres.K8sNodeTag{Id:nodeTagId, Key:req.Key , Value:req.Value, IsPublic:req.IsPublic}
	sta := postgres.K8sNodeTagAdd(nodeTag)
	if sta {
		resp_utils.ResponseOK(c, responseMsg.Success, "")
	} else {
		resp_utils.ResponseErrorParameter(c, "Failed To Add New Tag")
	}
}

func RootEditK8sNodeTag(ctx context.Context, c *app.RequestContext) {
	var req NodeTag
  err := c.BindAndValidate(&req)
	if err != nil {
		resp_utils.ResponseErrorParameter(c)
		return 
	}
	oldTag, err := postgres.K8sNodeTagInfo(req.Id)
	if err != nil {
		resp_utils.ResponseErrorParameter(c, "Error Tag Id")
		return
	}
	oldTag.Key			= req.Key
	oldTag.Value		= req.Value
	oldTag.IsPublic = req.IsPublic
	sta := postgres.K8sNodeTagUpdate(oldTag)
	if sta {
		resp_utils.ResponseOK(c, responseMsg.Success, "")
	} else {
		resp_utils.ResponseErrorParameter(c, "Failed To Edit Tag")
	}
}

func RootDelK8sNodeTag(ctx context.Context, c *app.RequestContext) {
	var req Id
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	sta := postgres.K8sNodeTagDel(req.Id)
	if sta {
		resp_utils.ResponseOK(c, responseMsg.Success, "")
	} else {
		resp_utils.ResponseErrorParameter(c, "Failed To Del The Tag")
	}
}

func RootGetResourcesList(ctx context.Context, c *app.RequestContext) {
	resourcesList, err := postgres.ResourcesListForRoot()
	if err != nil {
		resp_utils.ResponseError(c, "Get ResourcesList error:", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, resourcesList)
}

func RootAddResources(ctx context.Context, c *app.RequestContext) {
	type Reqbody struct {
		Value							map[string]interface{}		`json:"value,required"`
		IsPublic						bool										`json:"isPublic,required"`
	}
	var req Reqbody
  err := c.BindAndValidate(&req)
	if err != nil {
		resp_utils.ResponseErrorParameter(c)
		return 
	}
	resourcesId := resp_utils.RandStringWithLengthN(36)
	resources := &postgres.Resources{Id:resourcesId, Value:req.Value, IsPublic:req.IsPublic}
	sta := postgres.ResourcesAdd(resources)
	if !sta {
		resp_utils.ResponseError(c, "Add Resources error", nil)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, nil)
}

func RootEditResources(ctx context.Context, c *app.RequestContext) {
	type Reqbody struct {
		Id								string						`json:"id, require"`
		Value							map[string]interface{}		`json:"value"`
		IsPublic						bool						`json:"isPublic"`
	}
	var req Reqbody
    err := c.BindAndValidate(&req)
	if err != nil {
		resp_utils.ResponseErrorParameter(c)
		return 
	}
	resources := &postgres.Resources{Id:req.Id, Value:req.Value, IsPublic:req.IsPublic}
	sta := postgres.ResourcesUpdate(resources)
	if !sta {
		resp_utils.ResponseError(c, "Update Resources error", nil)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, nil)
}

func RootDelResources(ctx context.Context, c *app.RequestContext) {
	var req Id
	err := c.BindAndValidate(&req)
	if err != nil {
		resp_utils.ResponseError(c, "Get ResourcesList error:", err)
		return
	}
	sta := postgres.ResourcesDel(req.Id)
	if sta {
		resp_utils.ResponseOK(c, responseMsg.Success, "")
	} else {
		resp_utils.ResponseErrorParameter(c, "Failed To Del The Resources")
	}
}