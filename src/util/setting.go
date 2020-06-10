/*
@Time : 2020/6/10 10:26
@Author : wkang
@File : setting
@Description:
*/
package util

import (
	"github.com/go-ini/ini"
	"log"
)

var cfg *ini.File

type App struct {
	SourceRoot  string
	TargetFolder string
}

var AppSetting = &App{}

func init() {
	var err error
	cfg, err = ini.Load("src/config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config/app.ini': %v", err)
	}
	mapTo("app", AppSetting)


}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}