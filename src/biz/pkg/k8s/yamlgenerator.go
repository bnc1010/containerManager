package k8s

import (
	"fmt"
	// "encoding/json"
	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"github.com/bnc1010/containerManager/biz/utils"
	customError "github.com/bnc1010/containerManager/biz/pkg/error"
	"k8s.io/apimachinery/pkg/util/intstr"
)


func GenerateDeploymentYaml(deploymentName string, image string,portNums []interface{},replicas int32, k8snodetags map[string]interface{}, resources map[string]interface{}) (*appsv1.Deployment, error) {
	ports := []corev1.ContainerPort{}
	for _, _port := range portNums{
		ports = append(ports, corev1.ContainerPort{ContainerPort: utils.GetInterfaceToInt32(_port.(map[string] interface {})["port"]), Protocol: corev1.Protocol(_port.(map[string] interface {})["protocol"].(string))}) 
	}

	resourcesMM := *utils.MapII2MapMap(resources)

	_, hasLimits 			:= resourcesMM["limits"]
	_, hasRequests 			:= resourcesMM["requests"]
	limits					:= corev1.ResourceList{}
	requests 				:= corev1.ResourceList{}

	if hasLimits {
		limitsCPU 			:= resource.MustParse("100m")
		limitsMemory 		:= resource.MustParse("500Mi")
		_, hasLimitsCPU 	:= resourcesMM["limits"]["cpu"]
		if hasLimitsCPU {
			limitsCPU = resource.MustParse(resourcesMM["limits"]["cpu"])
		}
		_, hasLimitsMemory 	:= resourcesMM["limits"]["memory"]
		if hasLimitsMemory {
			limitsMemory = resource.MustParse(resourcesMM["limits"]["memory"])
		}
		limits = corev1.ResourceList{"cpu": limitsCPU,"memory": limitsMemory}
		for _k, _v := range resourcesMM["limits"] {
			if _k != "cpu" && _k != "memory" {
				q, err := resource.ParseQuantity(_v)
				if err == nil {
					limits[corev1.ResourceName(_k)] = q
				}
			}
		}
	}

	if hasRequests {
		requestsCPU 		:= resource.MustParse("100m")
		requestsMemory 		:= resource.MustParse("500Mi")
		_, hasRequestsCPU 	:= resourcesMM["requests"]["cpu"]
		if hasRequestsCPU {
			requestsCPU = resource.MustParse(resourcesMM["requests"]["cpu"])
		}
		_, hasRequestsMemory 	:= resourcesMM["requests"]["memory"]
		if hasRequestsMemory {
			requestsMemory = resource.MustParse(resourcesMM["requests"]["memory"])
		}
		requests = corev1.ResourceList{"cpu": requestsCPU,"memory": requestsMemory}
		for _k, _v := range resourcesMM["requests"] {
			if _k != "cpu" && _k != "memory"{
				q, err := resource.ParseQuantity(_v)
				if err == nil {
					requests[corev1.ResourceName(_k)] = q
				}
			}
		}
	}

	resourceRequirements := corev1.ResourceRequirements{Limits: limits, Requests: requests}
	fmt.Println(resourceRequirements)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":deploymentName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": deploymentName,
					},
				},
				Spec: corev1.PodSpec{
					NodeSelector: utils.MapInterface2String(k8snodetags),
					Containers: []corev1.Container{
						{
							Name:  deploymentName,
							Image: image,
							Ports: ports,
							Resources: resourceRequirements,
						},
					},
				},
			},
		},
	}
	return deployment, nil
}


func GenerateServiceYaml(serviceName string, selector map[string] string, portNums []interface{}) (*corev1.Service, error) {
	ports := []corev1.ServicePort {}
	for ind, _port := range portNums{
		usablePort := utils.RandUsablePort("47.100.69.138")
		if usablePort == -1 {
			return nil, &customError.CommonError{"no usable port now"}
		}
		ports = append(ports, corev1.ServicePort{Name:fmt.Sprintf("port-%d", ind), Port: utils.GetInterfaceToInt32(_port.(map[string] interface {})["port"]), TargetPort: intstr.FromInt(utils.GetInterfaceToInt(_port.(map[string] interface {})["port"])),Protocol: corev1.Protocol(_port.(map[string] interface {})["protocol"].(string)), NodePort: usablePort}) 
	}

	fmt.Println(ports)
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeNodePort,
			Selector: selector,
			Ports: ports,
		},
	}

	fmt.Println(service)
	return service, nil
}