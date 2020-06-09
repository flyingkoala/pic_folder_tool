package main

import (
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"os"
	"path/filepath"
)

var count = 0

var filetypes = []string{"JPG","HEIC","MOV","MP4","BMP","GIF","PNG"}

func logFileName(path string, info os.FileInfo, err error) error {

	count++

	// 返回错误后，编辑将终止
	//if count > 10 {
	//	return errors.New("stop")
	//}

	fmt.Printf("\n %d path: %s    fileName: %s     ", count, path, info.Name())
	return nil

}

func ReadOrientation(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("failed to open file, err: ", err)
		return 0
	}
	defer file.Close()

	x, err := exif.Decode(file)
	if err != nil {
		fmt.Println("failed to decode file, err: ", err)
		return 0
	}

	orientation, err := x.Get(exif.Orientation)
	if err != nil {
		fmt.Println("failed to get orientation, err: ", err)
		return 0
	}
	orientVal, err := orientation.Int(0)
	if err != nil {
		fmt.Println("failed to convert type of orientation, err: ", err)
		return 0
	}

	fmt.Println("the value of photo orientation is :", orientVal)
	return orientVal
}

func main() {
	filepath.Walk("D:\\图片1", logFileName)
}
