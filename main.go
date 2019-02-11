package main

import (
	"bee-myblog/cors"
	_ "bee-myblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	//if beego.BConfig.RunMode == "dev" {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/apidoc"] = "swagger"
	//}
	if err := beego.SetLogger("file", `{"filename":"logs/test.log"}`); err != nil {
		panic(err)
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))
	beego.Run()
}
