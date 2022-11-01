package k8s


import (
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


func CreateSimpleDeployment(namespaceName string,deploymentName string,image string,portNum int32,replicas int32)(deploymentInfo *appsv1.Deployment,err error)  {

	namespace := namespaceName
	//这个结构和原生k8s启动deployment的yml文件结构完全一样，对着写就好
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
					Containers: []corev1.Container{
						{
							Name:  deploymentName,
							Image: image,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: portNum,
								},
							},
						},
					},
				},
			},
		},
	}
	deploymentInfo,err = Client.AppsV1().Deployments(namespace).Create(context.TODO(),deployment,metav1.CreateOptions{})
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