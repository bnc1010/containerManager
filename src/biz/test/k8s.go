package test

import (
	"github.com/bnc1010/containerManager/biz/pkg/k8s"
)

func TestK8s() {
	k8s.PodHeapsterMemory("default", "51255903100-759-123-757595b856-gpjnz")
	k8s.PodHeapsterCpu("default", "51255903100-759-123-757595b856-gpjnz")
}