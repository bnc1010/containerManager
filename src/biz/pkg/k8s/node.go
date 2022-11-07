package k8s


import (
	"time"
	"context"
	"encoding/json"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NodeMetricsList struct {
    Kind       				  string	`json:"kind"`
    APIVersion 				  string 	`json:"apiVersion"`
    Metadata   struct {
        SelfLink 			  string 	`json:"selfLink"`
    } 									`json:"metadata"`
    Items []struct {
        Metadata struct {
            Name              string    `json:"name"`
            Namespace         string    `json:"namespace"`
            SelfLink          string    `json:"selfLink"`
            CreationTimestamp time.Time `json:"creationTimestamp"`
        } 								`json:"metadata"`
        Timestamp  time.Time 			`json:"timestamp"`
        Window     			  string    `json:"window"`
		Usage struct {
			Cpu 			  string 	`json:"cpu"`
			Memory 			  string 	`json:"memory"`
		} 								`json:usage`
    } 									`json:"items"`
}

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

func NodesMetrics() (nodes *NodeMetricsList, err error) {
	data, err := Client.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/nodes").DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &nodes)
	return nodes, nil
}

