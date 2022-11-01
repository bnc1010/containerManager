package k8s


import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNodeList()(nodeList *corev1.NodeList,err error)  {
	nodeList, err = Client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nodeList,err
	}
    return nodeList,nil
}


func GetNode(nodeName string)(nodeInfo *corev1.Node,err error)  {
	nodeInfo, err = Client.CoreV1().Nodes().Get(context.TODO(),nodeName,  metav1.GetOptions{})
	if err != nil {
		return nodeInfo, err
	}
	return nodeInfo,nil
}


