package postgres

import (
	"encoding/json"
	"github.com/bnc1010/containerManager/biz/utils"
)

func ImageInfo(imageId string) (*Image, error) {
	rows, err:= Client.Query("select * from tb_image where id=$1", imageId)
	defer rows.Close()
	if err!= nil{
		imageErrorLoger(err)
		return nil, err
	}
	var image * Image
	var bports []byte
	for rows.Next() {
		image = &Image{}
		err := rows.Scan(&image.Id, &image.Name, &image.Describe, &image.PullName, &image.Creator, &image.UseGPU,  &image.CreateTime, &image.UpdateTime, &image.Usable, &bports)
		if err != nil {
			imageErrorLoger(err)
			return nil, err
		}
		json.Unmarshal(bports, 	&image.Ports)
	}
	return image, nil
}

func ImageAdd(image *Image) bool {
	stmt, err := Client.Prepare("insert into tb_image(id,name,describe,pullname,creator,usegpu,usable,createtime,updatetime,ports) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)")
	defer stmt.Close()
	if err != nil {
		imageErrorLoger(err)
		return false
	}
	ports_json, _ := utils.Array2Bytes(image.Ports)
	_, err = stmt.Exec(image.Id, image.Name, image.Describe, image.PullName, image.Creator, image.UseGPU, image.Usable, image.CreateTime, image.UpdateTime, ports_json)
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

func ImageUpdate(image *Image) bool {
	stmt, err := Client.Prepare("update tb_image set name=$1,describe=$2,pullname=$3,usegpu=$4,usable=$5,updatetime=$6 where id=$7")
	defer stmt.Close()
	if err != nil {
		imageErrorLoger(err)
		return false
	}
	_, err = stmt.Exec(image.Name, image.Describe, image.PullName, image.UseGPU, image.Usable, image.UpdateTime, image.Id)
	if err != nil {
		imageErrorLoger(err)
		return false
	}
	return true
}

func ImagePublicCheck(imageId string) bool {
	image, err := ImageInfo(imageId)
	return !(err != nil || image == nil || !image.Usable)
}