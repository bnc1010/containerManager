package k8s


import (
	"fmt"
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func GetServiceList(namespaceName string) (serviceList *corev1.ServiceList,err error) {
	serviceList,err = Client.CoreV1().Services(namespaceName).List(context.TODO(),metav1.ListOptions{})
	if err != nil {
		return serviceList,err
	}
	return serviceList,nil
}

// apiVersion: v1
// kind: Service
// metadata:
//   name: nginx
//   namespace: test
//   labels:
//     app: nginx
// spec:
//   type: NodePort
//   ports:
//   - port: 80
//     targetPort: 80
//     nodePort: 30500
//   selector:
//     app: nginx
func CreateService(namespaceName string, serviceName string, selector map[string]string, portNums []interface{}) (serviceInfo *corev1.Service,err error) {
	service, err := GenerateServiceYaml(serviceName, selector, portNums)
	if err != nil {
		return nil, err
	}
	serviceInfo, err = Client.CoreV1().Services(namespaceName).Create(context.TODO(),service,metav1.CreateOptions{})
    if err != nil {
		fmt.Println(err)
	}
	return serviceInfo, nil
}

func GetService(namespaceName string,serviceName string) (serviceInfo *corev1.Service,err error) {
	serviceInfo,err = Client.CoreV1().Services(namespaceName).Get(context.TODO(),serviceName,metav1.GetOptions{})
	if err != nil {
		return serviceInfo,err
	}
	return serviceInfo,nil
}


func ApplyServiceByNodePort(namespaceName string,serviceName string,nodePort int32)(serviceInfo *corev1.Service,err error)  {

	service,err := Client.CoreV1().Services(namespaceName).Get(context.TODO(),serviceName,metav1.GetOptions{})
	if err !=nil {
		return serviceInfo,err
	}

    service.Spec.Ports[0].NodePort = nodePort
	serviceInfo,err = Client.CoreV1().Services(namespaceName).Update(context.TODO(),service,metav1.UpdateOptions{})
	if err !=nil {
		return serviceInfo,err
	}

	return serviceInfo,nil
}


func DeleteService(namespaceName string,serviceName string)(err error)  {
	err = Client.CoreV1().Services(namespaceName).Delete(context.TODO(),serviceName,metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}
