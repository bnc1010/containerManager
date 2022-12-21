package postgres

import (
	"encoding/json"
	"github.com/bnc1010/containerManager/biz/utils"
)


func (resources * Resources) Default() {
	resources.Value = map[string]interface{} {"limits":map[string]interface{} {"cpu":"500m", "memory":"1Gi"}}
	resources.IsPublic = true
} 


func ResourcesAdd(resources * Resources) bool {
	stmt, err := Client.Prepare("insert into tb_resource(id,value,ispublic) values($1,$2,$3)")
	defer stmt.Close()
	if err != nil {
		resourcesErrorLoger(err)
		return false
	}
	value_json, _ 		:= utils.Map2Bytes(resources.Value)
	_, err = stmt.Exec(resources.Id, value_json, resources.IsPublic)
	if err != nil {
		resourcesErrorLoger(err)
		return false
	}
	return true
}

func ResourcesInfo(resourcesId string)	(*Resources, error) {
	rows, err:= Client.Query("select * from tb_resources where id=$1", resourcesId)
	defer rows.Close()
	if err!= nil{
		resourcesErrorLoger(err)
		return nil, err
	}
	var resources * Resources
	var bvalue 			[]byte
	for rows.Next() {
		resources = & Resources{}
		err := rows.Scan(&resources.Id, &bvalue, &resources.IsPublic)
		if err != nil {
			resourcesErrorLoger(err)
			return nil, err
		}
		json.Unmarshal(bvalue, &resources.Value)
	}
	return resources, nil
}

func ResourcesForPublic() ([]*Resources, error)	{
	rows, err:= Client.Query("select * from tb_resources where ispublic=true")
	defer rows.Close()
	if err!= nil{
		resourcesErrorLoger(err)
		return nil, err
	}
	var resourcesList []*Resources
	var resources * Resources
	var bvalue 			[]byte
	for rows.Next() {
		resources = & Resources{}
		err := rows.Scan(&resources.Id, &bvalue, &resources.IsPublic)
		if err != nil {
			resourcesErrorLoger(err)
			return nil, err
		}
		json.Unmarshal(bvalue, &resources.Value)
		resourcesList = append(resourcesList, resources)
	}
	return resourcesList, nil
}

func ResourcesListForRoot() ([]*Resources, error)	{
	rows, err:= Client.Query("select * from tb_resources")
	defer rows.Close()
	if err!= nil{
		resourcesErrorLoger(err)
		return nil, err
	}
	var resourcesList []*Resources
	var resources * Resources
	var bvalue 			[]byte
	for rows.Next() {
		resources = & Resources{}
		err := rows.Scan(&resources.Id, &bvalue, &resources.IsPublic)
		if err != nil {
			resourcesErrorLoger(err)
			return nil, err
		}
		json.Unmarshal(bvalue, &resources.Value)
		resourcesList = append(resourcesList, resources)
	}
	return resourcesList, nil
}

func ResourcesUpdate(resources * Resources) bool {
	stmt, err := Client.Prepare("update tb_resources set value=$1,ispublic=$2 where id=$3")
	defer stmt.Close()
	if err != nil {
		resourcesErrorLoger(err)
		return false
	}
	value_json, _ 	:= utils.Map2Bytes(resources.Value)
	_, err = stmt.Exec(value_json, resources.IsPublic, resources.Id)
	if err != nil {
		resourcesErrorLoger(err)
		return false
	}
	return true
}