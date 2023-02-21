package handler


import (
	// "fmt"
	"github.com/bnc1010/containerManager/biz/pkg/postgres"
	"github.com/bnc1010/containerManager/biz/pkg/k8s"
)

func IsAdminOrRoot(role interface{}) bool {
	return role == "admin" || role == "root"
}

func IsRoot(role string) bool {
	return role == "root"
}


//
//	判断镜像是否属于project，且镜像没有被设置为禁用
//
func CheckImageBelongProject(imageId string, project * postgres.Project) (bool, *postgres.Image) {
	imageOK := false
	for _, _image := range project.Images {
		if _image == imageId {
			imageOK = true
			break
		}
	}
	if !imageOK {
		return false, nil
	}
	image, err := postgres.ImageInfo(imageId)
	// image在数据库中找不到
	if err != nil || image == nil {
		return false, nil
	}
	// image合法但是被设置禁用
	if !image.Usable {
		return false, nil
	} else {
		return true, image
	}
}


func CheckDeploymentExist(namespace string, deployment string) bool {
	deploymentInfo, err := k8s.GetDeployment(namespace, deployment)
	return deploymentInfo != nil && err ==nil
}