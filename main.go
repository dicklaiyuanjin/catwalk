package main

import (
	_ "catwalk/routers"
	"github.com/astaxie/beego"
)

func init() {
  beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}

