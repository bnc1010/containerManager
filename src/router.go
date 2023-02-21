// Code generated by hertz generator.

package main

import (
	handler "github.com/bnc1010/containerManager/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	r.POST("/postping", handler.PostTest)


	// common
	r.POST("/common/postping", handler.CommonPostTest)
	r.POST("/common/openProject", handler.CommonOpenProject)
	r.POST("/common/closeProject", handler.CommonCloseProject)
	r.POST("/common/getProjectUrl", handler.CommonGetProjectUrl)
	r.POST("/common/infoProject", handler.CommonProjectInfo)
	r.POST("/common/getProjectByUserId", handler.CommonProjectGetByUserId)
	r.POST("/common/createProject", handler.CommonCreateProject)
	
	
	// admin
	r.POST("/admin/postping", handler.AdminPostTest)
	// Pod
	r.POST("/admin/getPod", handler.AdminGetPod)
	r.POST("/admin/listPods", handler.AdminGetPodList)
	r.POST("/admin/podsMetrics", handler.AdminPodsMetrics)
	r.POST("/admin/podMemory", handler.AdminPodHeapsterMemory)
	r.POST("/admin/podCpu", handler.AdminPodHeapsterCpu)
	r.POST("/admin/podLog", handler.AdminPodLog)
	r.POST("/admin/podEvent", handler.AdminGetPodEvent)
	r.POST("/admin/listDeployments", handler.AdminGetDeploymentList)
	r.POST("/admin/getDeployment", handler.AdminGetDeployment)
	r.POST("/admin/listServices", handler.AdminGetServiceList)
	r.POST("/admin/listImages", handler.AdminGetImageList)

	r.POST("/admin/getFilesInfo", handler.AdminGetFilesInfo)
	r.POST("/admin/listFiles", handler.AdminGetFilesList)
	
	r.POST("/admin/getProjectInfo", handler.AdminGetProjectInfo)
	r.POST("/admin/listProjects", handler.AdminGetProjectsList)
	r.POST("/admin/listNodeTags", handler.AdminGetNodeTagsList)
	r.POST("/admin/listResources", handler.AdminGetResourcesList)
	
	r.POST("/admin/getImageInfo", handler.AdminGetImageInfo)
	r.POST("/admin/addImage", handler.AdminAddImage)
	r.POST("/admin/delImage", handler.AdminDelImage)
	r.POST("/admin/editImage", handler.AdminEditImage)

	
	r.POST("/admin/getDatasetInfo", handler.AdminGetDatasetInfo)
	r.POST("/admin/listDatasets", handler.AdminGetDatasetList)
	

	// root 
	r.POST("/root/checkPermission", handler.RootPermissionCheck)
	// Namespace
	r.POST("/root/getNamespace", handler.RootGetNamespace)
	r.POST("/root/listNamespaces", handler.RootGetNamespaceList)
	r.POST("/root/createNamespace", handler.RootCreateNamespace)
	r.POST("/root/deleteNamespace", handler.RootDeleteNamespace)
	// Node
	r.POST("/root/getNode", handler.RootGetNode)
	r.POST("/root/listNodes", handler.RootGetNodeList)
	r.POST("/root/nodesMetrics", handler.RootNodesMetrics)

	r.POST("/root/addNodeTag", handler.RootAddK8sNodeTag)
	r.POST("/root/delNodeTag", handler.RootDelK8sNodeTag)
	r.POST("/root/editNodeTag", handler.RootEditK8sNodeTag)

	r.POST("/root/editResources", handler.RootEditResources)
	r.POST("/root/addResources", handler.RootAddResources)
	r.POST("/root/delResources", handler.RootDelResources)

	
	
	
}
