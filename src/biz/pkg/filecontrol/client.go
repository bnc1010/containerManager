package filecontrol

import (
	"os"
	"fmt"
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