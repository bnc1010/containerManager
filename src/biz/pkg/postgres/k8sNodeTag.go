package postgres


func K8sNodeTagAdd(k8sNodeTag * K8sNodeTag) bool {
	stmt, err := Client.Prepare("insert into tb_k8snodetag(id,key,value,ispublic) values($1,$2,$3,$4)")
	defer stmt.Close()
	if err != nil {
		k8sNodeTagErrorLoger(err)
		return false
	}
	_, err = stmt.Exec(k8sNodeTag.Id, k8sNodeTag.Key, k8sNodeTag.Value, k8sNodeTag.IsPublic)
	if err != nil {
		k8sNodeTagErrorLoger(err)
		return false
	}
	return true
}

func K8sNodeTagInfo(k8sNodeTagId string)	(*K8sNodeTag, error) {
	rows, err:= Client.Query("select * from tb_k8snodetag where id=$1", k8sNodeTagId)
	defer rows.Close()
	if err!= nil{
		k8sNodeTagErrorLoger(err)
		return nil, err
	}
	var k8sNodeTag * K8sNodeTag
	for rows.Next() {
		k8sNodeTag = & K8sNodeTag{}
		err := rows.Scan(&k8sNodeTag.Id, &k8sNodeTag.Key, &k8sNodeTag.Value, &k8sNodeTag.IsPublic)
		if err != nil {
			k8sNodeTagErrorLoger(err)
			return nil, err
		}
	}
	return k8sNodeTag, nil
}

func K8sNodeTagList()	(*[]K8sNodeTag, error) {
	rows, err:= Client.Query("select * from tb_k8snodetag")
	defer rows.Close()
	if err!= nil{
		k8sNodeTagErrorLoger(err)
		return nil, err
	}
	var k8sNodeTags []K8sNodeTag
	for rows.Next() {
		k8sNodeTag := K8sNodeTag{}
		err := rows.Scan(&k8sNodeTag.Id, &k8sNodeTag.Key, &k8sNodeTag.Value, &k8sNodeTag.IsPublic)
		if err != nil {
			k8sNodeTagErrorLoger(err)
			continue
		}
		k8sNodeTags = append(k8sNodeTags, k8sNodeTag)
	}
	return &k8sNodeTags, nil
}

func K8sNodeTagForPublic() ([]*K8sNodeTag, error)	{
	rows, err:= Client.Query("select * from tb_k8snodetag where ispublic=true")
	defer rows.Close()
	if err!= nil{
		k8sNodeTagErrorLoger(err)
		return nil, err
	}
	var k8sNodeTagList []*K8sNodeTag
	var k8sNodeTag * K8sNodeTag
	for rows.Next() {
		k8sNodeTag = & K8sNodeTag{}
		err := rows.Scan(&k8sNodeTag.Id, &k8sNodeTag.Key, &k8sNodeTag.Value, &k8sNodeTag.IsPublic)
		if err != nil {
			k8sNodeTagErrorLoger(err)
			return nil, err
		}
		k8sNodeTagList = append(k8sNodeTagList, k8sNodeTag)
	}
	return k8sNodeTagList, nil
}