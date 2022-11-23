package postgres

func FilesInfo(filesId string) (*Files, error) {
	rows, err:= Client.Query("select * from tb_files where id=$1", filesId)
	defer rows.Close()
	if err!= nil{
		filesErrorLoger(err)
		return nil, err
	}
	var files * Files
	for rows.Next() {
		err := rows.Scan(&files.Id, &files.Name, &files.Creator, &files.Path, &files.CreateTime, &files.UpdateTime, &files.Size)
		if err != nil {
			filesErrorLoger(err)
			return nil, err
		}
	}
	return files, nil
}

func FilesAdd(files *Files) bool {
	stmt, err := Client.Prepare("insert into tb_files(id,name,creator,path,createtime,updatetime,size) values($1,$2,$3,$4,$5,$6,$7)")
	defer stmt.Close()
	if err != nil {
		filesErrorLoger(err)
		return false
	}
	_, err = stmt.Exec(files.Id, files.Name, files.Creator, files.Path, files.CreateTime, files.UpdateTime, files.Size)
	if err != nil {
		filesErrorLoger(err)
		return false
	}
	return true
}

func FilesDel(filesId string) bool {
	stmt, err := Client.Prepare("delete from tb_files where id=$1")
	defer stmt.Close()
	if err != nil {
		filesErrorLoger(err)
		return false
	}
	_, err = stmt.Exec(filesId)
	if err != nil {
		filesErrorLoger(err)
		return false
	}
	return true
}

func FilesUpdate(files * Files) bool {
	stmt, err := Client.Prepare("update tb_files set name=$1,path=$2,updatetime=$3 where id=$4")
	defer stmt.Close()
	if err != nil {
		filesErrorLoger(err)
		return false
	}
	_, err = stmt.Exec(files.Name, files.Path, files.UpdateTime, files.Id)
	if err != nil {
		filesErrorLoger(err)
		return false
	}
	return true
}