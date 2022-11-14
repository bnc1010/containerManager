package handler

type Namespace struct {
	Name string 		`json:"namespace,required"`
}

type Node struct {
	Name string 		`json:"node,required`
}

type Pod struct {
	Name string 		`json:"pod,required"`
	Namespace string 	`json:"namespace,required"`
}

type ResponseMsg struct {
	Success 		string
	SystemError 	string
	Forbidden		string
}

var responseMsg = ResponseMsg{"Success", "System Error", "No Permission"}