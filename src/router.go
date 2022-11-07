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

	
	// admin
	r.POST("/admin/postping", handler.AdminPostTest)
	// Pod
	r.POST("/admin/getPod", handler.AdminGetPod)
	r.POST("/admin/listPods", handler.AdminGetPodList)
	r.POST("/admin/podsMetrics", handler.AdminPodsMetrics)
	
	

	// root 
	r.POST("/root/postping", handler.RootPostTest)
	// Namespace
	r.POST("/root/getNamespace", handler.RootGetNamespace)
	r.POST("/root/listNamespaces", handler.RootGetNamespaceList)
	r.POST("/root/createNamespace", handler.RootCreateNamespace)
	r.POST("/root/deleteNamespace", handler.RootDeleteNamespace)
	// Node
	r.POST("/root/getNode", handler.RootGetNode)
	r.POST("/root/listNodes", handler.RootGetNodeList)
	r.POST("/root/nodesMetrics", handler.RootNodesMetrics)
	
}
