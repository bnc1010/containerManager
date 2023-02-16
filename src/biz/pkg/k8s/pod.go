package k8s


import (
    "io"
    "fmt"
	"time"
    "math"
    "bytes"
	"context"
    "strings"
	"encoding/json"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	customError "github.com/bnc1010/containerManager/biz/pkg/error"
)



func ParseRFC3339(s string, nowFn func() metav1.Time) (metav1.Time, error) {
    if t, timeErr := time.Parse(time.RFC3339Nano, s); timeErr == nil {
        return metav1.Time{Time: t}, nil
    }
    t, err := time.Parse(time.RFC3339, s)
    if err != nil {
        return metav1.Time{}, err
    }
    return metav1.Time{Time: t}, nil
}

func TimeNow() metav1.Time {
    return metav1.Time{time.Now()}
}

func SetlogOption(containerName string, follow bool, timestamps bool, sincetime *string, since *string, tailLines  *int64) (*corev1.PodLogOptions, error) {
    logOptions              := corev1.PodLogOptions{}
    logOptions.Follow       = follow
    logOptions.Container    = containerName
    logOptions.Timestamps   = timestamps
    var sinceSeconds time.Duration
    var err error
    //sincetime和since同时只能设置一个
    if sincetime != nil && since != nil {
        return nil, &customError.CommonError{"can not set since and sincetime both"}
    }
    //如果都没有设置，就设置默认10分钟
    if sincetime == nil && since == nil {
        sinceSeconds, err = time.ParseDuration("10m")
        if err != nil {
            return nil, err
        }
    }
    //如果since设置，则解析
    if since != nil {
        sinceSeconds, err = time.ParseDuration(*since)
        if err != nil {
            return nil, err
        }
        if sinceSeconds > 0 {
            // round up to the nearest second
            sec := int64(math.Ceil(float64(sinceSeconds) / float64(time.Second)))
            logOptions.SinceSeconds = &sec
        }
    }
    //如果sincetime设置，则默认+8小时
    if sincetime != nil {
        if !strings.Contains(*sincetime, "+") {
            *sincetime = *sincetime + "+08:00"
        }
        if ts, err := ParseRFC3339(*sincetime, TimeNow); err == nil {
            logOptions.SinceTime = &ts
        }
    }
    //赋值获取行数，默认是全部
    if *tailLines > 0 {
        logOptions.TailLines = tailLines
    }
    return &logOptions, nil
}



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

func GetPodList(namespaceName string, fieldSelector string, labelSelector string) (podList *corev1.PodList, err error) {
	podList, err = Client.CoreV1().Pods(namespaceName).List(context.TODO(), metav1.ListOptions{FieldSelector: fieldSelector, LabelSelector: labelSelector})
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

func PodsMetrics() (pods *PodMetricsList, err error) {
	data, err := Client.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/pods").DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &pods)
	return pods, nil
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

func GetContainersOfPod(namespaceName string, podName string, containerName * string) ([]string, error) {
    podContainersNames := []string{}
    podInfo, err := GetPod(namespaceName, podName)
    if err != nil {
        return podContainersNames, err
    }
    for _, container := range podInfo.Spec.Containers {
        fmt.Println(container)
        if containerName == nil{
            podContainersNames = append(podContainersNames, container.Name)
            continue
        }
        if container.Name == *containerName {
            podContainersNames = append(podContainersNames, container.Name)
            break
        }
    }
    return podContainersNames, nil
}

func PodLog(namespaceName string,podName string, containerName * string, sincetime *string, since *string, tailLines *int64) (string, error) {
    // fmt.Println(podName, sincetime, since, tailLines)
    containers, err := GetContainersOfPod(namespaceName, podName, containerName)
    if err != nil {
        return "", err
    }
    if len(containers) < 1 {
        return "", &customError.CommonError{"no containers in this pod"}
    }
    logOptions, err := SetlogOption(containers[0], false, false, sincetime, since, tailLines)
    if err != nil {
        return "", err
    }
    // fmt.Println(logOptions)
    stream, err := Client.CoreV1().Pods(namespaceName).GetLogs(podName, logOptions).Stream(context.TODO())
    defer stream.Close()
    buf := new(bytes.Buffer)
    _, err = io.Copy(buf, stream)
    if err != nil {
        return "", err
    }
    str := buf.String()
    return str, nil
}