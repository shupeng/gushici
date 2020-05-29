package main

import (
	"gushici/models"
	_ "gushici/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.SetStaticPath("/uploads", "uploads")
	beego.Run()
}
