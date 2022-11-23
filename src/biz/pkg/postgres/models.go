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
}

type Image struct {
	Id				string
	Name			string
	Describe		string
	PullName		string
	Creator			string
	UseGPU			bool
	CreateTime		time.Time
	UpdateTime		time.Time
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














