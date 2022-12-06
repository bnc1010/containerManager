package test

import (
	"fmt"
	"time"
	"github.com/bnc1010/containerManager/biz/utils"
	"github.com/bnc1010/containerManager/biz/pkg/postgres"
)

func TestPostgres() {
	// testImageAdd()
	// testDatasetAdd()
	// testFilesAdd()
	testProjectAdd()
}

// Id				string
// Name			string
// Describe		string
// PullName		string
// Creator			string
// UseGPU			bool
// CreateTime		time.Time
// UpdateTime		time.Time

func testImageAdd() {
	var t_now = time.Now()
	var count_s = 0
	var count_f = 0
	for ind :=1; ind<=10;ind ++ {
		var image = postgres.Image{Id:"thisisarandstrforid" + utils.RandStringWithLengthN(5) , Name:"sampleimagename", Describe:"some msg for describe", PullName:"bnc1010/xxxx:v0.1", Creator:"thisisarandstrforid", UseGPU:false, Usable:true, CreateTime:t_now, UpdateTime:t_now}
		sta := postgres.ImageAdd(&image)
		if sta {
			count_s += 1
			// fmt.Println("success")
		} else {
			count_f += 1
			// fmt.Println("fail")
		}
	}
	var t_two = time.Now()
	fmt.Println("success: ", count_s, " fail: ", count_f)
	fmt.Println(t_now, t_two)
}

// Creator			string
// CreateTime		time.Time
// UpdateTime		time.Time
// Path			string
// IsPublic		bool
// Size			int64
func testDatasetAdd() {
	var t_now = time.Now()
	var count_s = 0
	var count_f = 0
	for ind :=1; ind<=10;ind ++ {
		var dataset = postgres.Dataset{Id:"thisisarandstrforid" + utils.RandStringWithLengthN(5) , Name:"sampledatasetname", Describe:"some msg for describe",  Creator:"thisisarandstrforid",  CreateTime:t_now, UpdateTime:t_now, Path:"../pathsample", IsPublic: true, Size:1000}
		sta := postgres.DatasetAdd(&dataset)
		if sta {
			count_s += 1
			// fmt.Println("success")
		} else {
			count_f += 1
			// fmt.Println("fail")
		}
	}
	var t_two = time.Now()
	fmt.Println("success: ", count_s, " fail: ", count_f)
	fmt.Println(t_now, t_two)
}


// Id				string
// Name			string
// Creator			string
// Path			string
// CreateTime		time.Time
// UpdateTime		time.Time
// Size			int64
func testFilesAdd() {
	var t_now = time.Now()
	var count_s = 0
	var count_f = 0
	for ind :=1; ind<=10;ind ++ {
		var files = postgres.Files{Id:"thisisarandstrforid" + utils.RandStringWithLengthN(5) , Name:"samplefilesname",  Creator:"thisisarandstrforid",  CreateTime:t_now, UpdateTime:t_now, Path:"../pathsample", Size:1000}
		sta := postgres.FilesAdd(&files)
		if sta {
			count_s += 1
			// fmt.Println("success")
		} else {
			count_f += 1
			// fmt.Println("fail")
		}
	}
	var t_two = time.Now()
	fmt.Println("success: ", count_s, " fail: ", count_f)
	fmt.Println(t_now, t_two)
}


// Id				string
// Name			string
// Describe		string
// Owner			string
// CreateTime		time.Time
// LastOpenTime	time.Time
// IsPublic		bool
// Files			map[string]string
// Datasets		map[string]string
// Images			[] string
// ForkFrom		string
// K8sNodeTags		map[string]string
// Resources		map[string]string
func testProjectAdd() {
	var t_now = time.Now()
	var count_s = 0
	var count_f = 0
	for ind :=1; ind<=3;ind ++ {
		var project = postgres.Project{Id:"thisisarandstrforid" + utils.RandStringWithLengthN(5) , Name:"sampleprojectname", Describe:"describe info",  Owner:"thisisarandstrforid",  CreateTime:t_now, LastOpenTime:t_now, IsPublic:false , Files:map[string] interface{} {"thisisarandstrforidlhuOJ":"/test"}, Datasets:map[string] interface{} {"thisisarandstrforidlhuOJ":"/test"}, Images:[] interface{} {"i1","i2"}, ForkFrom:"fork from sample", K8sNodeTags:map[string]interface{}{"cal_type":"cpu"}, Resources:map[string]interface{} {"limits":map[string]interface{} {"cpu":"500m", "memory":"1Gi"}}}
		sta := postgres.ProjectAdd(&project)
		if sta {
			count_s += 1
			// fmt.Println("success")
		} else {
			count_f += 1
			// fmt.Println("fail")
		}
	}
	var t_two = time.Now()
	fmt.Println("success: ", count_s, " fail: ", count_f)
	fmt.Println(t_now, t_two)
}