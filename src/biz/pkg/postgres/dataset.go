package postgres

func DatasetInfo(datasetId string) (*Dataset, error) {
	rows, err:= Client.Query("select * from tb_dataset where id=$1", datasetId)
	defer rows.Close()
	if err!= nil{
		datasetErrorLoger(err)
		return nil, err
	}
	var dataset * Dataset
	for rows.Next() {
		err := rows.Scan(&dataset.Id, &dataset.Name, &dataset.Describe, &dataset.Creator, &dataset.Path, &dataset.CreateTime, &dataset.UpdateTime, &dataset.Size, &dataset.IsPublic)
		if err != nil {
			datasetErrorLoger(err)
			return nil, err
		}
	}
	return dataset, nil
}

func DatasetAdd(dataset * Dataset) bool {
	stmt, err := Client.Prepare("insert into tb_dataset(id,name,describe,creator,path,createtime,updatetime,size,ispublic) values($1,$2,$3,$4,$5,$6,$7,$8,$9)")
	defer stmt.Close()
	if err != nil {
		datasetErrorLoger(err)
		return false
	}
	_, err = stmt.Exec(dataset.Id, dataset.Name, dataset.Describe, dataset.Creator, dataset.Path, dataset.CreateTime, dataset.UpdateTime, dataset.Size, dataset.IsPublic)
	if err != nil {
		datasetErrorLoger(err)
		return false
	}
	return true
}

func DatasetDel(datasetId string) bool {
	stmt, err := Client.Prepare("delete from tb_dataset where id=$1")
	defer stmt.Close()
	if err != nil {
		datasetErrorLoger(err)
		return false
	}
	_, err = stmt.Exec(datasetId)
	if err != nil {
		datasetErrorLoger(err)
		return false
	}
	return true
}

func DatasetUpdate(dataset * Dataset) bool {
	stmt, err := Client.Prepare("update tb_dataset set name=$1,describe=$2,updatetime=$3,ispublic=$4 where id=$5")
	defer stmt.Close()
	if err != nil {
		datasetErrorLoger(err)
		return false
	}
	_, err = stmt.Exec(dataset.Name, dataset.Describe, dataset.UpdateTime, dataset.IsPublic)
	if err != nil {
		datasetErrorLoger(err)
		return false
	}
	return true
}