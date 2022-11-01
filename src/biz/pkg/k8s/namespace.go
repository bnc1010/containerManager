package k8s


import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNamespaceList()(namespaceList *corev1.NamespaceList,err error)  {
	namespaceList,err = Client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return namespaceList,err
	}
	return namespaceList,nil
}

func GetNamespace(namespaceName string)(namespaceInfo *corev1.Namespace,err error)  {
	namespaceInfo,err = Client.CoreV1().Namespaces().Get(context.TODO(),namespaceName,metav1.GetOptions{})
	if err != nil {
		return namespaceInfo,err
	}
	return namespaceInfo,nil
}

func CreateNamespace(namespaceName string)(namespaceInfo *corev1.Namespace,err error)  {
	var namespace corev1.Namespace
	namespace.Name = namespaceName
	namespaceInfo,err = Client.CoreV1().Namespaces().Create(context.TODO(),&namespace,metav1.CreateOptions{})
	if err != nil {
		return namespaceInfo,err
	}
	return namespaceInfo,nil
}

func DeleteNamespace(namespaceName string)(err error)  {
	err = Client.CoreV1().Namespaces().Delete(context.TODO(),namespaceName,metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}