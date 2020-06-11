package main

import (
	"fmt"
	"github.com/src/util"
	"path/filepath"
)




func main() {
	md5FilePath:=fmt.Sprintf("%s\\md5.txt",util.AppSetting.TargetFolder)

	fmt.Println("源文件的根目录为：  ",util.AppSetting.SourceRoot)
	util.PathCreate(util.AppSetting.TargetFolder)//目标文件夹不存在则创建
	util.FileCreate(md5FilePath) //MD5文件不存在则创建
	fmt.Println("目标文件的目录为：  ",util.AppSetting.TargetFolder)
	err:= util.ReadMD5Txt(md5FilePath)
	if err!=nil{
		return
	}
	filepath.Walk(util.AppSetting.SourceRoot, util.DealFile)
	util.WriteMD5Txt()
}
