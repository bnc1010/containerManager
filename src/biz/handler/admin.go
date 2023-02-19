package handler

import (
	"fmt"
	"time"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/bnc1010/containerManager/biz/pkg/k8s"
	"github.com/bnc1010/containerManager/biz/pkg/postgres"
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
	podList,err := k8s.GetPodList(req.Name, req.FieldSelector, req.LabelSelector)
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

func AdminGetPodEvent(ctx context.Context, c *app.RequestContext) {
	var req Namespace
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	eventList, err := k8s.GetEventList(req.Name, req.FieldSelector)
	if err != nil {
		resp_utils.ResponseError(c, "Get Pod Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, eventList)
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
	logs, err := k8s.PodLog(req.Namespace, req.Pod, &req.Container, nil, &sinceTime, &req.TailLines)
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

func AdminGetDeploymentList(ctx context.Context, c *app.RequestContext) {
	var req Namespace
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	deploymentList,err := k8s.GetDeploymentList(req.Name, req.FieldSelector)
	if err != nil {
		resp_utils.ResponseError(c, "Get DeploymentList Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, deploymentList)
}

func AdminGetDeployment(ctx context.Context, c *app.RequestContext) {
	var req Deployment
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	deploymentInfo, err := k8s.GetDeployment(req.Namespace, req.Name)
	if err != nil {
		resp_utils.ResponseError(c, "Get Pod Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, deploymentInfo)
}


func AdminGetServiceList(ctx context.Context, c *app.RequestContext) {
	var req Namespace
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	serviceList,err := k8s.GetServiceList(req.Name, req.FieldSelector, req.LabelSelector)
	if err != nil {
		resp_utils.ResponseError(c, "Get ServiceList Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, serviceList)
}


func AdminGetImageInfo(ctx context.Context, c *app.RequestContext) {
	var req Id
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	image, err := postgres.ImageInfo(req.Id)
	if err != nil {
		resp_utils.ResponseError(c, "Get ImageInfo Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, image)
}


func AdminGetImageList(ctx context.Context, c *app.RequestContext) {
	images, err := postgres.ImageList()
	if err != nil {
		resp_utils.ResponseError(c, "Get ImageList Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, images)
}

func AdminAddImage(ctx context.Context, c *app.RequestContext) {
	requestUserId := fmt.Sprintf("%v",ctx.Value("requestUserId"))
	var req ImageView
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	nowTime := time.Now()
	imageId := resp_utils.RandStringWithLengthN(36)
	image := postgres.Image{Id:imageId, Name:req.Name, Describe:req.Describe, PullName:req.PullName, Creator:requestUserId, UseGPU:req.UseGPU, Usable:req.Usable, CreateTime:nowTime, UpdateTime:nowTime, Ports:[]interface{}{}}
	
	for _, port := range req.Ports {
		image.Ports = append(image.Ports, port)
	}
	sta := postgres.ImageAdd(&image)
	if sta {
		resp_utils.ResponseOK(c, responseMsg.Success, "")
	} else {
		resp_utils.ResponseErrorParameter(c, "Failed To Add The Image")
	}
}

func AdminDelImage(ctx context.Context, c *app.RequestContext) {
	var req Id
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	sta := postgres.ImageDel(req.Id)
	if sta {
		resp_utils.ResponseOK(c, responseMsg.Success, "")
	} else {
		resp_utils.ResponseErrorParameter(c, "Failed To Del The Image")
	}
}

func AdminEditImage(ctx context.Context, c *app.RequestContext) {
	var req ImageView
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	nowTime := time.Now()
	oldImage, err := postgres.ImageInfo(req.Id)
	if err != nil {
		resp_utils.ResponseErrorParameter(c, "Error Image Id")
		return
	}
	oldImage.Name = req.Name
	oldImage.Describe = req.Describe
	oldImage.PullName = req.PullName
	oldImage.UseGPU = req.UseGPU
	oldImage.Usable = req.Usable
	oldImage.UpdateTime = nowTime
	oldImage.Ports = []interface{}{}
	
	for _, port := range req.Ports {
		oldImage.Ports = append(oldImage.Ports, port)
	}
	sta := postgres.ImageUpdate(oldImage)
	if sta {
		resp_utils.ResponseOK(c, responseMsg.Success, "")
	} else {
		resp_utils.ResponseErrorParameter(c, "Failed To Add The Image")
	}
}


func AdminGetFilesInfo(ctx context.Context, c *app.RequestContext) {
	var req Id
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	files, err := postgres.FilesInfo(req.Id)
	if err != nil {
		resp_utils.ResponseError(c, "Get FileInfo Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, files)
}


func AdminGetFilesList(ctx context.Context, c *app.RequestContext) {
	filesList, err := postgres.FilesList()
	if err != nil {
		resp_utils.ResponseError(c, "Get FilesList Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, filesList)
}

func AdminGetDatasetInfo(ctx context.Context, c *app.RequestContext) {
	var req Id
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	dataset, err := postgres.DatasetInfo(req.Id)
	if err != nil {
		resp_utils.ResponseError(c, "Get DatasetInfo Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, dataset)
}

func AdminGetDatasetList(ctx context.Context, c *app.RequestContext) {
	datasetList, err := postgres.DatasetList()
	if err != nil {
		resp_utils.ResponseError(c, "Get DatasetList Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, datasetList)
}

func AdminGetProjectInfo(ctx context.Context, c *app.RequestContext) {
	var req Id
	err := c.BindAndValidate(&req)
	if err != nil {
		hlog.CtxErrorf(ctx, "[PostTest] Unmarshal failed, err: %v", err)
		resp_utils.ResponseErrorParameter(c)
		return
	}
	project, err := postgres.ProjectInfo(req.Id)
	if err != nil {
		resp_utils.ResponseError(c, "Get ProjectInfo Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, project)
}

func AdminGetProjectsList(ctx context.Context, c *app.RequestContext) {
	filesList, err := postgres.ProjectList()
	if err != nil {
		resp_utils.ResponseError(c, "Get ProjectList Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, filesList)
}

func AdminGetNodeTagsList(ctx context.Context, c *app.RequestContext) {
	nodeTagsList, err := postgres.K8sNodeTagList()
	if err != nil {
		resp_utils.ResponseError(c, "Get ProjectList Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, nodeTagsList)
}

func AdminGetResourcesList(ctx context.Context, c *app.RequestContext) {
	resourcesList, err := postgres.ResourcesList()
	if err != nil {
		resp_utils.ResponseError(c, "Get ResourcesList Error", err)
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, resourcesList)
}






