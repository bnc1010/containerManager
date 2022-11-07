package handler

type Namespace struct {
	Name string 		`json:"name,required"`
}

type Node struct {
	Name string 		`json:"name,required`
}

type Pod struct {
	Name string 		`json:"name,required"`
	Namespace string 	`json:"namespace,required"`
}