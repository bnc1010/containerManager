package k8s


import (
	"fmt"
	"context"
	corev1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func GetDeploymentList(namespaceName string)(deploymentList *appsv1.DeploymentList,err error)  {
	deploymentList,err = Client.AppsV1().Deployments(namespaceName).List(context.TODO(), metav1.ListOptions{})
	if err != nil{
		return nil, err
	}
	return deploymentList, err
}

func GetDeployment(namespaceName string,deploymentName string)(deploymentInfo *appsv1.Deployment,err error)  {
	deploymentInfo,err = Client.AppsV1().Deployments(namespaceName).Get(context.TODO(), deploymentName,metav1.GetOptions{})
	if err != nil{
		return nil, err
	}
	return deploymentInfo, err
}

func GetPodsOfDeployment(namespaceName string,deploymentName string) (podList *corev1.PodList, err error){
	deploymentInfo,err := Client.AppsV1().Deployments(namespaceName).Get(context.TODO(), deploymentName,metav1.GetOptions{})
	if err != nil{
		return nil, err
	}
	selector, err := metav1.LabelSelectorAsSelector(deploymentInfo.Spec.Selector)
	if err != nil {
		return nil, err
	}
	listOptions := metav1.ListOptions{LabelSelector: selector.String()}
	podList, err = Client.CoreV1().Pods(namespaceName).List(context.TODO(), listOptions)
	return podList, nil
}

func CreateSimpleDeployment(namespaceName string,deploymentName string,image string,ports []interface{},replicas int32, k8snodetags map[string]interface{}, resources map[string]interface{})(deploymentInfo *appsv1.Deployment,err error)  {
	deployment, err := GenerateDeploymentYaml(deploymentName,image, ports, replicas, k8snodetags, resources)
	if err != nil {
		return nil, err
	}
	// fmt.Println(deployment)
	deploymentInfo,err = Client.AppsV1().Deployments(namespaceName).Create(context.TODO(),deployment,metav1.CreateOptions{})
    if err != nil {
		return deploymentInfo,err
	}
	return deploymentInfo,nil
}

func ApplyDeploymentByImage(namespaceName string,deploymentName string,image string)(deploymentInfo *appsv1.Deployment,err error)  {
	deployment,err := Client.AppsV1().Deployments(namespaceName).Get(context.TODO(),deploymentName,metav1.GetOptions{})
	if err !=nil {
		return deploymentInfo,err
	}
	deployment.Spec.Template.Spec.Containers[0].Image = image
	deploymentInfo,err = Client.AppsV1().Deployments(namespaceName).Update(context.TODO(),deployment,metav1.UpdateOptions{})
	if err !=nil {
		return deploymentInfo,err
	}
	return deploymentInfo,nil
}

func ApplyDeploymentByReplicas(namespaceName string,deploymentName string,replicas int32)(deploymentInfo *appsv1.Deployment,err error)  {

	deployment,err := Client.AppsV1().Deployments(namespaceName).Get(context.TODO(),deploymentName,metav1.GetOptions{})
	if err !=nil {
		return deploymentInfo,err
	}

	deployment.Spec.Replicas = &replicas
	deploymentInfo,err = Client.AppsV1().Deployments(namespaceName).Update(context.TODO(),deployment,metav1.UpdateOptions{})
	if err !=nil {
		return deploymentInfo,err
	}
	return deploymentInfo,nil
}

func DeleteDeployment(namespaceName string,deploymentName string)(err error)  {
	err = Client.AppsV1().Deployments(namespaceName).Delete(context.TODO(),deploymentName,metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}


func GetPodsLogOfDeployment(namespaceName string,deploymentName string) (map[string]string, error) {
	mp := map[string] string {}
	podList, err := GetPodsOfDeployment(namespaceName, deploymentName)
	if err != nil {
		return mp, err
	}
	
	tailLines := int64(100)
	since := "72h"
	for _, item := range podList.Items {
		log, err := PodLog(namespaceName ,item.ObjectMeta.Name, nil, &since, &tailLines)
		if err == nil {
			mp[item.ObjectMeta.Name] = log
		} else {
			fmt.Println(err)
		}
	}
	return mp, nil
}