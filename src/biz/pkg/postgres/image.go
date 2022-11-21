package postgres

func ImageInfo(imageId string) (*Image, error) {
	rows, err:= Client.Query("select * from tb_image where id=$1", imageId)
	defer rows.Close()
	if err!= nil{
		imageErrorLoger(err)
		return nil, err
	}
	var image * Image
	for rows.Next() {
		err := rows.Scan(&image.Id, &image.Name, &image.Describe, &image.PullName, &image.Creator, &image.UseGPU, &image.CreateTime, &image.UpdateTime)
		if err != nil {
			imageErrorLoger(err)
			return nil, err
		}
	}
	return &image, nil
}

func ImageAdd(image *Image) bool {
	stmt, err := Client.Prepare("insert into tb_image(id,name,describe,pullname,creator,usegpu,createtime,updatetime) values($1,$2,$3,$4,$5,$6,$7,$8)")
	defer stmt.Close()
	if err != nil {
		imageErrorLoger(err)
		return false
	}
	_, err = stmt.Exec(image.Id, image.Name, image.Describe, image.PullName, image.Creator, image.UseGPU, image.CreateTime, image.UpdateTime)
	if err != nil {
		imageErrorLoger(err)
		return false
	}
	return true
}

func ImageDel(imageId string) bool {
	stmt, err := Client.Prepare("delete from tb_image where id=$1")
	defer stmt.Close()
	if err != nil {
		imageErrorLoger(err)
		return false
	}
	_, err = stmt.Exec(imageId)
	if err != nil {
		imageErrorLoger(err)
		return false
	}
	return true
}

func ImageUpdate(image *Image) {
	stmt, err := Client.Prepare("update tb_image set name=$1,describe=$2,pullname=$3,usegpu=$4,updatetime=$5 where id=$6")
	defer stmt.Close()
	if err != nil {
		imageErrorLoger(err)
		return false
	}
	_, err = stmt.Exec(image.Name, image.Describe, image.PullName, image.UseGPU, image.UpdateTime, image.Id)
	if err != nil {
		imageErrorLoger(err)
		return false
	}
	return true
}