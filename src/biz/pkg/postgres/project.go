package postgres

import (
	"fmt"
	"encoding/json"
	"github.com/bnc1010/containerManager/biz/utils"
)

func (project * Project) Mask() {
	project.Files 		= nil
	project.Images		= nil
	project.Datasets 	= nil
	project.ForkFrom 	= ""	
	project.Resources 	= nil
	project.K8sNodeTags = nil
}



func ProjectInfo(projectId string)	(*Project, error) {
	fmt.Println(projectId)
	rows, err:= Client.Query("select * from tb_project where id=$1", projectId)
	defer rows.Close()
	if err!= nil{
		projectErrorLoger(err)
		return nil, err
	}
	var project * Project
	var bfiles 			[]byte
	var bdatasets 		[]byte
	var bimages 		[]byte
	var bk8snodeTags 	[]byte
	var bresources 		[]byte
	for rows.Next() {
		project = & Project{}
		err := rows.Scan(&project.Id, &project.Name, &project.Describe, &project.Owner, &project.CreateTime, &project.LastOpenTime, &project.IsPublic, &bfiles, &bdatasets, &bimages, &project.ForkFrom, &bk8snodeTags, &bresources)
		if err != nil {
			projectErrorLoger(err)
			return nil, err
		}
	}
	json.Unmarshal(bfiles, 		&project.Files)
	json.Unmarshal(bdatasets, 	&project.Datasets)
	json.Unmarshal(bimages, 	&project.Images)
	json.Unmarshal(bk8snodeTags,&project.K8sNodeTags)
	json.Unmarshal(bresources, 	&project.Resources)
	return project, nil
}

func ProjectAdd(project *Project) bool {
	stmt, err := Client.Prepare("insert into tb_project(id,name,describe,owner,createtime,lastopentime,ispublic,files,datasets,images,forkfrom,k8snodetags,resources) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)")
	defer stmt.Close()
	if err != nil {
		projectErrorLoger(err)
		return false
	}
	files_json, _ 		:= utils.Map2Bytes(project.Files)
	images_json, _ 		:= utils.Array2Bytes(project.Images)
	datasets_json, _ 	:= utils.Map2Bytes(project.Datasets)
	resources_json, _ 	:= utils.Map2Bytes(project.Resources)
	k8snodetags_json, _ := utils.Map2Bytes(project.K8sNodeTags)
	_, err = stmt.Exec(project.Id, project.Name, project.Describe, project.Owner, project.CreateTime, project.LastOpenTime, project.IsPublic, files_json, datasets_json, images_json, project.ForkFrom, k8snodetags_json, resources_json)
	if err != nil {
		projectErrorLoger(err)
		return false
	}
	return true
}

func ProjectDel(projectId string) bool {
	stmt, err := Client.Prepare("delete from tb_project where id=$1")
	defer stmt.Close()
	if err != nil {
		projectErrorLoger(err)
		return false
	}
	_, err = stmt.Exec(projectId)
	if err != nil {
		projectErrorLoger(err)
		return false
	}
	return true
}

func ProjectUpdate(project *Project) bool {
	stmt, err := Client.Prepare("update tb_project set name=$1,describe=$2,lastopentime=$3,ispublic=$4,files=$5,datasets=$6,images=$7 where id=$8")
	defer stmt.Close()
	if err != nil {
		projectErrorLoger(err)
		return false
	}
	_, err = stmt.Exec(project.Name, project.Describe, project.LastOpenTime, project.IsPublic, project.Files, project.Datasets, project.Images, project.Id)
	if err != nil {
		projectErrorLoger(err)
		return false
	}
	return true
}
