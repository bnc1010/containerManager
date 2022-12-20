package filecontrol

import (
	"os"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/bnc1010/containerManager/biz/utils"
)

type Path struct {
	Root			string
	UserFilePath 	string
	DatasetPath		string
}

var SystemPath = Path{Root:"", UserFilePath:"", DatasetPath:""}

func InitPath() bool {
	ctx := context.Background()

	//----------单独配置两个子目录的情况-------------------
	if len(os.Getenv("FILE_UserFilePath")) != 0 {
		SystemPath.UserFilePath = os.Getenv("FILE_UserFilePath")
	}
	if len(os.Getenv("DatasetPath")) != 0 {
		SystemPath.DatasetPath = os.Getenv("FILE_DatasetPath")
	}
	//----------------------------------------------------

	// 没有配置两个子目录，也没有配置ROOT则直接报错
	if len(os.Getenv("DatasetPath")) == 0 && len(os.Getenv("FILE_UserFilePath")) == 0 && len(os.Getenv("FILE_ROOT")) == 0 {
		hlog.CtxErrorf(ctx, "[FilePath] FILE_ROOT Haven't Set")
		return false
	} else {
		SystemPath.Root = os.Getenv("FILE_ROOT")
	}



	//----如果没有单独配置两个子目录，则在ROOT下创建默认------
	if len(SystemPath.UserFilePath) == 0 {
		SystemPath.UserFilePath = SystemPath.Root + "/userfile"
	}
	if len(SystemPath.DatasetPath) == 0 {
		SystemPath.DatasetPath = SystemPath.Root + "/dataset"
	}
	//-----------------------------------------------------

	if len(SystemPath.Root) != 0 {
		exist, _  := utils.PathExists(SystemPath.Root)
		if !exist {
			err := os.Mkdir(SystemPath.Root, 0755)
			if err != nil {
				hlog.CtxErrorf(ctx, "[FilePath] FILE_ROOT has set, but not exist and create failed")
				return false
			}
		}
	}

	exist, _  := utils.PathExists(SystemPath.UserFilePath)
	if !exist {
		err := os.Mkdir(SystemPath.UserFilePath, 0755)
		if err != nil {
			hlog.CtxErrorf(ctx, "[FilePath] FILE_UserFilePath create failed")
			return false
		}
	}

	exist, _  = utils.PathExists(SystemPath.DatasetPath)
	if !exist {
		err := os.Mkdir(SystemPath.DatasetPath, 0755)
		if err != nil {
			hlog.CtxErrorf(ctx, "[FilePath] FILE_DatasetPath create failed")
			return false
		}
	}

	hlog.CtxInfof(ctx, "[FilePath] FILE_ROOT: %#v",SystemPath.Root)
	return true
}

func GenerateFilePath(userId string, fileId string) (string, error) {
	userPath := SystemPath.UserFilePath + "/" + userId
	exist, _  := utils.PathExists(userPath)
	if !exist {
		err := os.Mkdir(userPath, 0755)
		if err != nil {
			return "", err
		}
	}
	
	filePath := userPath + "/" + fileId
	exist, _  = utils.PathExists(filePath)
	if !exist {
		err := os.Mkdir(filePath, 0755)
		if err != nil {
			return "", err
		}
	}
	return filePath, nil
}

func GenerateDatasetpath(userId string, datasetId string) (string, error) {
	userPath := SystemPath.DatasetPath + "/" + userId
	exist, _  := utils.PathExists(userPath)
	if !exist {
		err := os.Mkdir(userPath, 0755)
		if err != nil {
			return "", err
		}
	}

	datasetPath := userPath + "/" + datasetId
	exist, _  = utils.PathExists(datasetPath)
	if !exist {
		err := os.Mkdir(datasetPath, 0755)
		if err != nil {
			return "", err
		}
	}
	return datasetPath, nil
}


func scanDir(path string, syncM *sync.Map, wait *sync.WaitGroup){
	defer wait.Done()
	dirAry,err := ioutil.ReadDir(path)
	if err != nil {
		panic(err);
	}
	fmt.Println(dirAry)
	for _,e := range dirAry{
		if e.IsDir(){
			wait.Add(1)
			go scanDir(filepath.Join(path,e.Name()), syncM, wait)
		}else{
			syncM.Store(filepath.Join(path,e.Name()),(e.Size()))
		}
	}
}	

func CalDirSize(path string) (int64, int64) {
	fmt.Println(path)
	var syncM sync.Map
	var wait sync.WaitGroup
	wait.Add(1)
	go scanDir(path, &syncM, &wait)
	wait.Wait()
	var fileCount int64
	var dirSize int64
	syncM.Range(func(key, value interface{}) bool {
		fileCount++
		v := value.(int64)
		fmt.Println(v)
		dirSize += v
		return true
	})
	return dirSize, fileCount
}
