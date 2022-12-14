package postgres

import (
	"time"
)

type Dataset struct {
	Id				string
	Name			string
	Describe		string
	Creator			string
	CreateTime		time.Time
	UpdateTime		time.Time
	Path			string
	IsPublic		bool
	Size			int64
}

type Project struct {
	Id				string
	Name			string
	Describe		string
	Owner			string
	CreateTime		time.Time
	LastOpenTime	time.Time
	IsPublic		bool
	Files			map[string]interface{}
	Datasets		map[string]interface{}
	Images			[]interface{}
	ForkFrom		string


	// 下面两项只有root可以修改，否则都是创建时从系统中选择
	K8sNodeTags		map[string]interface{}
	Resources		map[string]interface{}
	Usable			bool
}

type Image struct {
	Id				string
	Name			string
	Describe		string
	PullName		string
	Creator			string
	UseGPU			bool
	Usable			bool
	CreateTime		time.Time			
	UpdateTime		time.Time
	Ports			[]interface{}
}

type Files struct {
	Id				string
	Name			string
	Creator			string
	Path			string
	CreateTime		time.Time
	UpdateTime		time.Time
	Size			int64
}


type Resources struct {
	Id				string
	Value			map[string]interface{}
	IsPublic		bool
}


type K8sNodeTag struct {
	Id				string
	Key				string
	Value			string
	IsPublic		bool
}