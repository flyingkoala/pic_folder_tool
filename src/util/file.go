/*
@Time : 2020/6/10 10:49
@Author : wkang
@File : file
@Description:
*/
package util

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"os"
	p "path"
	"strings"
)

var count =0
//处理文件
func DealFile(path string, info os.FileInfo, err error) error {

	//log.Println(AppSetting.FileTypes)
	filetype:= strings.ToLower(p.Ext(p.Base(path))) //获取文件的后缀名
	if filetype!=""{
		photoyear:=ReadPhotoYear(path)
		if photoyear!=""{
			count++
			tfolder:=fmt.Sprintf("%s\\%s",AppSetting.TargetFolder,photoyear)
			PathCreate(tfolder)
			fmt.Println(fmt.Sprintf("序号：%d,拍摄年份：%s, 图片路径：%s,目标路径：%s",count,photoyear,path,tfolder))
		}

	}

	//fmt.Printf("\n %d path: %s    fileName: %s     ", count, path, info.Name())
	return nil

}

//读取图片extif信息中的拍摄时间
func ReadPhotoYear(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	x, err := exif.Decode(file)
	if err != nil {
		return ""
	}

	shottime, err := x.Get(exif.DateTime)
	if err != nil {
		return ""
	}

	ts:= strings.Split(shottime.String(),":")
	if len(ts)==0{
		return ""
	}
	return ts[0][1:5]
}

// 判断文件夹是否存在，不存在则创建
func PathCreate(path string)   {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		}
	}
}
