package handler

type Namespace struct {
	Name 					string 					`json:"namespace,required"`
	FieldSelector string 					`json:"fieldSelector"`
	LabelSelector string 					`json:"labelSelector"`
}

type Node struct {
	Name string 									`json:"node,required`
}

type Pod struct {
	Name 					string 					`json:"pod,required"`
	Namespace 		string 					`json:"namespace,required"`
}

type Deployment struct {
	Name 					string 					`json:"deployment,required"`
	Namespace 		string 					`json:"namespace,required"`
}

type PodLog struct {		
	Pod string 										`json:"pod,required"`
	Namespace string 							`json:"namespace,required"`
	TailLines int64								`json:"lines,required"`
	Container string							`json:"container,required"`
}

func (podlog * PodLog) VaildTailLines() bool {
	if podlog.TailLines > 0 && podlog.TailLines < 10000 && (podlog.TailLines % 50) == 0 {
		return true
	} else {
		return false
	}
}

type ResponseMsg struct {
	Success 		string
	SystemError 	string
	Forbidden		string
}

var responseMsg = ResponseMsg{"Success", "System Error", "No Permission"}
var sinceTime string = "8640h"