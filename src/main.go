package main

import (
	"fmt"
	"github.com/src/util"
	"path/filepath"
)




func main() {
	fmt.Println("源文件的根目录为：  ",util.AppSetting.SourceRoot)
	util.PathCreate(util.AppSetting.TargetFolder)//创建目标文件夹
	fmt.Println("目标文件的目录为：  ",util.AppSetting.TargetFolder)

	filepath.Walk(util.AppSetting.SourceRoot, util.DealFile)
}
