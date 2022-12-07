package test

import (
	"fmt"
	"github.com/bnc1010/containerManager/biz/pkg/k8s"
)

func TestK8s() {
	// data_1, err_1 := k8s.PodHeapsterMemory("default", "51255903100-759-123-757595b856-gpjnz")
	// if err_1 == nil {
	// 	fmt.Println(data_1)
	// }
	// data_2, err_2 := k8s.PodHeapsterCpu("default", "51255903100-759-123-757595b856-gpjnz")
	// if err_2 == nil {
	// 	fmt.Println(data_2)
	// }
	mp, err := k8s.GetPodsLogOfDeployment("default", "10205501424-121-124")
	fmt.Println(mp, err)
}