package handler

import (
	"fmt"
	"time"
	"regexp"
	"strings"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/bnc1010/containerManager/biz/pkg/postgres"
	"github.com/bnc1010/containerManager/biz/pkg/k8s"
	"github.com/bnc1010/containerManager/biz/pkg/filecontrol"
	resp_utils "github.com/bnc1010/containerManager/biz/utils"
)

var (
	regestr = `token=[0-9a-z]{48}`
	reg = regexp.MustCompile(regestr)
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
	requestUserId := ctx.Value("requestUserId")
	requestUserRole := ctx.Value("requestUserRole")
	type Reqbody struct {
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
	if !IsAdminOrRoot(requestUserRole) && project.Owner != requestUserId {
		resp_utils.ResponseForbid(c, fmt.Sprintf("have no perimission of the project"))
		return 
	}

	//
	// todo: 动态配置的namespace，以及deploymentName怎么给，每个用户最多同时可以开多少个
	//
	namespace				:= "default"
	deploymentName	:= project.Owner + "-" + strings.ToLower(req.ProjectId)
	serviceName			:= "service-" + deploymentName

	// 如果已经存在对应dployment，直接返回信息
	isExist := CheckDeploymentExist(namespace, deploymentName)
	if isExist {
		resp_utils.ResponseOK(c, responseMsg.Success, deploymentName)
		return
	}

	// 检查image是否是该project拥有的
	imageSta, image := CheckImageBelongProject(req.ImageId, project)
	if !imageSta {
		resp_utils.ResponseErrorParameter(c, "image illegal")
		return
	}

	// 尝试创建deployment
	_, err = k8s.CreateSimpleDeployment(namespace, deploymentName, image.PullName, image.Ports, 1, project.K8sNodeTags, project.Resources)
	if err != nil {
		resp_utils.ResponseError(c, "some thing error when open the deployment", err)
		return
	}
	_, err = k8s.CreateService(namespace, serviceName, map[string]string {"app": deploymentName}, image.Ports)
	if err != nil {
		resp_utils.ResponseError(c, "some thing error when create the service", err)
		k8s.DeleteDeployment(namespace, deploymentName)
		return 
	}
	resp_utils.ResponseOK(c, responseMsg.Success, deploymentName)
}
//curl -d '{"userId":"423h4huhuhfuseu34", "projectId":"thisisarandstrforidqQgma", "imageId":"thisisarandstrforidPOwjG"}' -H "Content-Type:application/json" -H "AUTH_TOKEN:Aa2N9jIOFz4If8Qn/EPGAn2nTd4z0BkcM45E6YetcGI1x9NOgDkUQFftPcNaAI6R "  -X POST http://127.0.0.1:8888/common/openProject

func CommonGetProjectUrl(ctx context.Context, c *app.RequestContext){
	requestUserId := ctx.Value("requestUserId")
	requestUserRole := ctx.Value("requestUserRole")
	type Reqbody struct {
		Deployment				string `json:"deployment,required"`
	}
	req := Reqbody{}
	err := c.BindAndValidate(&req)
	if err != nil {
		resp_utils.ResponseErrorParameter(c)
		return 
	}

	//检查deployment的拥有者是否是当前用户
	if requestUserId != strings.Split(req.Deployment, "-")[0] && !IsAdminOrRoot(requestUserRole) {
		resp_utils.ResponseForbid(c, fmt.Sprintf("have no perimission of the project"))
		return 
	}

	namespace := "default"
	isExist := CheckDeploymentExist(namespace, req.Deployment)
	if ! isExist {
		resp_utils.ResponseOK(c, responseMsg.Success, "not exist")
		return
	}
	base_url := "http://139.224.216.129:"
	access_url := base_url
	service, err := k8s.GetService(namespace, "service-" + req.Deployment)

	ports := service.Spec.Ports
	for _, _port := range ports {
		if _port.NodePort > 0{
			access_url = access_url + fmt.Sprintf("%d", _port.NodePort)
			break
		}
	}
	logs, err := k8s.GetPodsLogOfDeployment(namespace, req.Deployment)
	token := ""
	for _, _v := range(logs) {
		resdata := reg.FindAllStringSubmatch(_v,-1)
		if len(resdata) > 0 {
			token = resdata[len(resdata)-1][0][6:]
			break
		}
	}
	type Resbody struct {
		Url			string
		Token		string
	}
	if len(token) == 0{
		resp_utils.ResponseOK(c, responseMsg.Success, "not ready")
		return
	}
	resp_utils.ResponseOK(c, responseMsg.Success, Resbody{access_url, token})
}

func CommonCloseProject(ctx context.Context, c *app.RequestContext) {
	requestUserId := ctx.Value("requestUserId")
	requestUserRole := ctx.Value("requestUserRole")
	type Reqbody struct {
		Deployment				string `json:"deployment,required"`
	}
	req := Reqbody{}
	err := c.BindAndValidate(&req)
	if err != nil {
		resp_utils.ResponseErrorParameter(c)
		return 
	}
	//检查deployment的拥有者是否是当前用户
	if requestUserId != strings.Split(req.Deployment, "-")[0] && !IsAdminOrRoot(requestUserRole) {
		resp_utils.ResponseForbid(c, fmt.Sprintf("have no perimission of the project"))
		return 
	}
	namespace := "default"
	
	k8s.DeleteDeployment(namespace, req.Deployment)
	k8s.DeleteService(namespace, "service-" + req.Deployment)
	resp_utils.ResponseOK(c, responseMsg.Success, "success")
}

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
		project.Mask()
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
	// 需要检查这些project是否是当前用户拥有的&其他用户公开的
	for _,project := range projects {
		if project.Owner == requestUserId || project.IsPublic {
			resProjects = append(resProjects, project)
		}
	}
	resp_utils.ResponseOK(c, responseMsg.Success, resProjects)
}


func CommonCreateProject(ctx context.Context, c *app.RequestContext) {
	requestUserId := fmt.Sprintf("%v",ctx.Value("requestUserId"))
	type Reqbody struct {
		ProjectName 					string	`json:"projectName,required"`
		Describe						string	`json:"describe"`
		IsPublic						bool	`json:"isPublic"`
		Images						  []string	`json:"images"`
		K8sNodeTagIds				  []string	`json:"k8sNodeTagIds"`
		ResourcesId						string	`json:"resourcesId"`
	}
	var req Reqbody
    err := c.BindAndValidate(&req)
	if err != nil {
		resp_utils.ResponseErrorParameter(c)
		return 
	}
	nowTime := time.Now()
	projectId := resp_utils.RandStringWithLengthN(36)
	project := postgres.Project{Id:projectId, Name:req.ProjectName, Describe:req.Describe, Owner:requestUserId, CreateTime:nowTime, IsPublic:req.IsPublic, Usable:true, Files:map[string]interface{}{}, Datasets:map[string]interface{}{}, Images:[]interface{}{}, K8sNodeTags:map[string]interface{}{}, Resources:map[string]interface{}{}}
	fmt.Println(project)
	// 创建专属文件File
	// fileId := resp_utils.RandStringWithLengthN(36)
	fileId := "JWFzbHYEejRSSJZjqogRAmzbTWieHlJQbtLk"
	newFilePath, err := filecontrol.GenerateFilePath(requestUserId, fileId)
	if err != nil {
		resp_utils.ResponseError(c, "create project error", err)
		return
	}
	dirSize,dirCount := filecontrol.CalDirSize(newFilePath)
	file := postgres.Files{Id:fileId , Name:req.ProjectName + "-file", Creator:requestUserId, Path:newFilePath, CreateTime:nowTime, UpdateTime:nowTime, Size:dirSize}
	fmt.Println(dirSize,dirCount,file)
	// sta := postgres.FilesAdd(&file)
	// if !sta {
	// 	resp_utils.ResponseError(c, "create project error", err)
	// 	return
	// }
	project.Files[fileId] = "/userfile"
	// 绑定Images
	for _, imageId := range req.Images {
		if postgres.ImagePublicCheck(imageId) {
			project.Images = append(project.Images, imageId)
		}
	}
	// 绑定Datasets

	// K8sNodeTags
	var k8sNodeTagList []* postgres.K8sNodeTag
	for _, t := range req.K8sNodeTagIds {
		k8sNodeTag, err := postgres.K8sNodeTagInfo(t)
		if k8sNodeTag != nil && err == nil {
			k8sNodeTagList = append(k8sNodeTagList, k8sNodeTag)
		}
	}
	project.FillK8sNodeTags(k8sNodeTagList)
	// Resources
	resources, err := postgres.ResourcesInfo(req.ResourcesId)
	if resources == nil || err != nil || !resources.IsPublic {
		resources = &postgres.Resources{}
		resources.Default()
	}
	fmt.Println(resources)
	project.FillResources(resources)
	fmt.Println(project)
	sta := postgres.ProjectAdd(&project)
	if sta {
		resp_utils.ResponseOK(c, responseMsg.Success, "")
	} else {
		resp_utils.ResponseErrorParameter(c, "Failed To Create The Project")
	}
}
//curl -d '{"projectName":"testproject","describe":"something for describe","isPublic":true, "images":["thisisarandstrforidPOwjG"]}' -H "Content-Type:application/json" -H "AUTH_TOKEN:Aa2N9jIOFz4If8Qn/EPGAn2nTd4z0BkcM45E6YetcGI1x9NOgDkUQFftPcNaAI6R"  -X POST http://127.0.0.1:8888/common/createProject


// func CommonForkProject(ctx context.Context, c *app.RequestContext) {

// }