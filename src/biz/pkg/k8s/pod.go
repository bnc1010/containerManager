package k8s


import (
	"time"
	"context"
	"encoding/json"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodMetricsList struct {
    Kind       string `json:"kind"`
    APIVersion string `json:"apiVersion"`
    Metadata   struct {
        SelfLink string `json:"selfLink"`
    } `json:"metadata"`
    Items []struct {
        Metadata struct {
            Name              string    `json:"name"`
            Namespace         string    `json:"namespace"`
            SelfLink          string    `json:"selfLink"`
            CreationTimestamp time.Time `json:"creationTimestamp"`
        } `json:"metadata"`
        Timestamp  time.Time `json:"timestamp"`
        Window     string    `json:"window"`
        Containers []struct {
            Name  string `json:"name"`
            Usage struct {
                CPU    string `json:"cpu"`
                Memory string `json:"memory"`
            } `json:"usage"`
        } `json:"containers"`
    } `json:"items"`
}

type PodUsage struct {
    Metrics [] struct {
        Timestamp  time.Time    `json:"timestamp"`
        Value      int64       `json:"value"`
    }                           `json:"metrics"`
    LatestTimestamp time.Time   `json:"latestTimestamp"`
}

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

func PodsMetrics() (nodes *PodMetricsList, err error) {
	data, err := Client.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/pods").DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &nodes)
	return nodes, nil
}

func PodHeapsterMemory(namespaceName string,podName string)  (podUsage * PodUsage, err error)  {
    var path = "/api/v1/namespaces/kube-system/services/heapster/proxy/api/v1/model/namespaces/" + namespaceName + "/pods/" + podName + "/metrics/memory/usage"
    podHeapster, err := Client.RESTClient().Get().AbsPath(path).DoRaw(context.TODO())
    if err != nil {
        return nil, err
	}
    json.Unmarshal(podHeapster, &podUsage)
    return podUsage, nil
}

func PodHeapsterCpu(namespaceName string,podName string) (podUsage * PodUsage, err error) {
    var path = "/api/v1/namespaces/kube-system/services/heapster/proxy/api/v1/model/namespaces/" + namespaceName + "/pods/" + podName + "/metrics/cpu/usage_rate"
    podHeapster, err := Client.RESTClient().Get().AbsPath(path).DoRaw(context.TODO())
    if err != nil {
        return nil, err
	}
    json.Unmarshal(podHeapster, &podUsage)
    return podUsage, nil
}