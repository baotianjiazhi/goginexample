package main

import (
	"ginexample/models"
	"ginexample/router"
	"ginexample/setting"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	r := router.Newrouter()
	r.Run(setting.ServerSetting.HttpPort)
}
