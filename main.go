package main

import (
	"MVP/pkg/setting"
	"log"
)

func main() {
	log.Println("Hello, api 正在启动中...")
	setting.SetUp()                             //初始化配置文件
	log.Println(setting.ServerSetting.HttpPort) //测试能否打印出ini配置文件设置的信息
}
