package controllers

import (
	"github.com/astaxie/beego"
  "catwalk/models"
  "html/template"
)

type AppController struct {
	beego.Controller
}

func (this *AppController) App() {
  usr := this.GetSession("username").(string)
  userinfo := models.UserInfoJSON{Username: usr}
  var icon template.URL
  if models.ReadUserInfo(&userinfo, "username") == true {
    if userinfo.Icon == "" {
      icon = template.URL("/static/img/icon.png")
    } else {
      icon = template.URL(userinfo.Icon)
    }
  }
  this.Data["userinfo"] = &userinfo
  this.Data["icon"] = icon
  this.TplName = "app.tpl"
}

func (this *AppController) AppSignout() {
  //注销session
  //修改数据库数据表user中相应的isactive字段
  u := this.GetSession("username")
  this.DelSession("username")
  b := models.SignErr{State: 0}
  if models.SetUserUnActive(u.(string)) == true {
    b.State = 1
  }
  this.Data["json"] = b
  this.ServeJSON()

}

func (this *AppController) AppEdit() {
  b := models.UserinfoErr{
    Existnick: 1,
    Existemail: 1,
  }
  var userinfo models.UserInfoJSON
  resbody := this.Ctx.Input.RequestBody
  models.EditHelper(&userinfo, resbody, &b)
  this.Data["json"] = b
  this.ServeJSON()
}

func (this *AppController) AppUpload() {
  b := models.SignErr{State: 0}
  var userinfo models.UserInfoJSON
  resbody := this.Ctx.Input.RequestBody
  if models.AnalyzeUserInfoJson(&userinfo, resbody) == true {
    if models.UpdateIconOfUserInfo(&userinfo) == true {
      b.State = 1
    }
  }
  this.Data["json"] = b
  this.ServeJSON()

}
