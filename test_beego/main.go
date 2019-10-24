package main

import (
	"github.com/astaxie/beego"
	_ "hello/test_beego/routers"
)

func main() {
	// 注册静态文件目录
	beego.SetStaticPath("/download", "download")
	beego.SetStaticPath("/upload", "upload")

	beego.Run()
}
