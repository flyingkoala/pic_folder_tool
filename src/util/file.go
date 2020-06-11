/*
@Time : 2020/6/10 10:49
@Author : wkang
@File : file
@Description:
*/
package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"io"
	"io/ioutil"
	"os"
	p "path"
	"strings"
)

var MD5Map = make(map[string]int)
var count =0
//处理文件
func DealFile(path string, info os.FileInfo, err error) error {

	//log.Println(AppSetting.FileTypes)
	filetype:= strings.ToLower(p.Ext(p.Base(path))) //获取文件的后缀名
	if filetype!=""{
		photoyear:=ReadPhotoYear(path)
		if photoyear==""{
			photoyear=ReadPhotoYearBasic(path)
		}
		if photoyear!=""{
			count++
			tfolder:=fmt.Sprintf("%s\\%s",AppSetting.TargetFolder,photoyear)
			PathCreate(tfolder)

			md5str,err:= GetFileMD5(path)
			if err==nil{
				//移动文件到对应目录
				if MD5Map[md5str]!=1{
					newpath:=tfolder+"\\"+info.Name()
					err:=os.Rename(path,newpath)
					if err!=nil{
						fmt.Println("移动失败",err.Error())
					}
					fmt.Println(fmt.Sprintf("序号：%d, md5：%s， 拍摄年份：%s, 图片路径：%s, 目标路径：%s",count,md5str,photoyear,path,tfolder))
				}else {
					fmt.Println("文件已存在目标路径")
				}
				//将MD5值写入全局map
				MD5Map[md5str]=1


			}


		}

	}
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
	fmt.Println(shottime)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	fmt.Println(shottime.String())
	ts:= strings.Split(shottime.String(),":")
	if len(ts)==0{
		return ""
	}
	if ts[0]=="\"\""{
		return ""
	}

	return ts[0][1:5]
}

//读取图片信息中的创建时间
func ReadPhotoYearBasic(filename string) string {
	fileInfo, _ := os.Stat(filename)
	if fileInfo==nil{
		return ""
	}
	//修改时间
	modTime := fileInfo.ModTime().Format("2006-01-02 15:04:05")
	ts:= strings.Split(modTime,"-")
	if len(ts)==0||fileInfo.IsDir(){
		return ""
	}
	return ts[0]


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
// 判断文件是否存在，不存在则创建
func FileCreate(filepath string)   {
	file,err:=os.Open(filepath)
	defer file.Close()
	if err!=nil && os.IsNotExist(err) {
		file, _ = os.Create(filepath)

	}
}
//读取MD5文本文件内容，将文件中的数据存入全局MD5map中
func ReadMD5Txt(filepath string)  error  {
	file,err:=os.Open(filepath)
	defer file.Close()
	if err!=nil  {
		return err
	}
	//1、一次性读取文件内容,还有一个 ReadAll的函数，也能读取
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	md5all:=string(data)
	md5arr:= strings.Split(md5all,",")
	for _,v:=range md5arr{
		if v!=""{
			MD5Map[v]=1
		}
	}

	return nil
}
//将图片的md5追加写入到MD5文件中
func WriteMD5Txt(){
	//将全局md5值存入md5文件中
	md5s:=""
	for k,_:=range MD5Map{
		md5s=md5s+","+k
	}
	file, _ := os.OpenFile(AppSetting.TargetFolder+"\\md5.txt",  os.O_WRONLY | os.O_CREATE | os.O_APPEND , 0666)
	// 写入文件内容
	io.WriteString(file, md5s)

}

//获取文件的md5值
func GetFileMD5(filepath string) (string,error) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err == nil {
		md5h := md5.New()
		io.Copy(md5h, file)
		MD5Str := hex.EncodeToString(md5h.Sum(nil))
		return MD5Str,nil
	}
	return "",err
}