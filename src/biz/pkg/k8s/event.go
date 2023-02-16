package k8s


import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetEventList(namespaceName string, fieldSelector string) (eventList *corev1.EventList, err error) {
	eventList, err = Client.CoreV1().Events(namespaceName).List(context.TODO(), metav1.ListOptions{FieldSelector: fieldSelector})
	if err != nil {
		return eventList, err
	}
	return eventList, nil
}



