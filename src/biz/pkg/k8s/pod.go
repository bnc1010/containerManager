package k8s


import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func GetPodList(namespaceName string) (podList *corev1.PodList, err error) {
	podList, err = Client.CoreV1().Pods(namespaceName).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return podList, err
	}
	return podList, nil
}

func GetPod(namespaceName string,podName string) (podInfo *corev1.Pod, err error) {
	podInfo, err = Client.CoreV1().Pods(namespaceName).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return podInfo, err
	}
	return podInfo, nil
}