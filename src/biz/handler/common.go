package handler

import (
	"fmt"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/bnc1010/containerManager/biz/pkg/postgres"
	"github.com/bnc1010/containerManager/biz/pkg/k8s"
	resp_utils "github.com/bnc1010/containerManager/biz/utils"
)

func CommonPostTest(ctx context.Context, c *app.RequestContext) {
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



func CommonOpenProject(ctx context.Context, c *app.RequestContext) {
	type Reqbody struct {
		UserId 					string `json:"userId,required"`
		ProjectId				string `json:"projectId,required"`
		ImageId					string `json:"imageId,required"`
	}
	var req Reqbody
    err := c.BindAndValidate(&req)
	if err != nil {
		resp_utils.ResponseErrorParameter(c)
		return 
	}
	project, err := postgres.ProjectInfo(req.ProjectId)
	if err != nil {
		resp_utils.ResponseErrorParameter(c)
		return
	}
	// 检查project是否是当前用户的
	if project.Owner != req.UserId {
		resp_utils.ResponseForbid(c, fmt.Sprintf("have no perimission of the project"))
		return 
	}
	// 检查image是否是该project拥有的
	imageOK := false
	if len(req.ImageId) != 0 {
		for _, _image := range project.Images {
			if _image == req.ImageId {
				imageOK = true
				break
			}
		}
	}
	var image * postgres.Image
	if imageOK {
		image, err = postgres.ImageInfo(req.ImageId)
		fmt.Println(image)
		if err != nil || image == nil {
			imageOK = false
		}
		// image合法但是被设置禁用
		if !image.Usable {
			resp_utils.ResponseErrorParameter(c, "chosen image is be set to unusable")
			return
		}
	}
	if !imageOK {
		resp_utils.ResponseErrorParameter(c, "no usable image")
		return
	}

	namespace			:= "default"
	deploymentName		:= req.UserId + "-1"
	serviceName			:= "service-" + deploymentName

	// 尝试创建deployment
	_, err = k8s.CreateSimpleDeployment(namespace, deploymentName, image.PullName, image.Ports, 1, project.K8sNodeTags, project.Resources)
	if err != nil {
		resp_utils.ResponseError(c, "some thing error when open the deployment", err)
	}
	_, err = k8s.CreateService(namespace, serviceName, map[string]string {"app": deploymentName}, image.Ports)
	if err != nil {
		resp_utils.ResponseError(c, "some thing error when create the service", err)
		k8s.DeleteDeployment(namespace, deploymentName)
	}
	resp_utils.ResponseOK(c, responseMsg.Success, deploymentName)
}
//curl -d '{"userId":"423h4huhuhfuseu34", "projectId":"thisisarandstrforidqQgmb", "imageId":"thisisarandstrforidPOwjG"}' -H "Content-Type:application/json" -H "AUTH_TOKEN:Aa2N9jIOFz4If8Qn/EPGAn2nTd4z0BkcM45E6YetcGI1x9NOgDkUQFftPcNaAI6R "  -X POST http://127.0.0.1:8888/common/openProject


func CommonProjectInfo(ctx context.Context, c *app.RequestContext) {
	type Reqbody struct {
		UserId 					string `json:"userId,required"`
		ProjectId				string `json:"projectId,required"`
	}
	var req Reqbody
    err := c.BindAndValidate(&req)
	if err != nil {
		resp_utils.ResponseErrorParameter(c)
		return 
	}
	project, err := postgres.ProjectInfo(req.ProjectId)
	if err != nil {
		resp_utils.ResponseErrorParameter(c)
		return
	}
	if project.Owner != req.UserId && project.IsPublic == false {
		fmt.Println("debug")
		project.Mask()
		fmt.Println(project)
	}
	resp_utils.ResponseOK(c, responseMsg.Success, project)
}

func CommonProjectGetByUserId(ctx context.Context, c *app.RequestContext) {
	requestUserId := ctx.Value("requestUserId")
	type Reqbody struct {
		UserId 					string `json:"userId,required"`
	}
	var req Reqbody
    err := c.BindAndValidate(&req)
	if err != nil {
		resp_utils.ResponseErrorParameter(c)
		return 
	}
	projects, err := postgres.ProjectsGetByUserId(req.UserId)
	if err != nil {
		resp_utils.ResponseError(c, "get projects error", err)
		return
	}
	
	resProjects  := [] *postgres.Project {}
	//需要检查这些project是否是当前用户拥有的&其他用户公开的
	for _,project := range projects {
		if project.Owner == requestUserId || project.IsPublic {
			resProjects = append(resProjects, project)
		}
	}
	resp_utils.ResponseOK(c, responseMsg.Success, resProjects)
}