package controllers

import (
	"github.com/astaxie/beego"
  "catwalk/models"
  "html/template"
  "fmt"
)

type AppController struct {
	beego.Controller
}

func (this *AppController) App() {
  usr := this.GetSession("username").(string)

  //setting part
  var userinfo models.JsUif
  userinfo.Username = usr
  var icon template.URL
  if models.Crud.Uif.Read(&userinfo, "username") == true {
    if userinfo.Icon == "" {
      icon = template.URL("/static/img/icon.png")
    } else {
      icon = template.URL(userinfo.Icon)
    }
  }
  this.Data["userinfo"] = userinfo
  this.Data["icon"] = icon

  //invitation part
  //作为reciver，获得所有sender发送给自己的invitation
  var inviArr []models.JsIvtt
  ok := models.Crud.Invitation.ReadList(&inviArr, userinfo.Username, "Receiver")
  fmt.Println(inviArr)
  if ok {
    this.Data["Invitations"] = inviArr
  }

  this.TplName = "app.tpl"
}

/*****************************************
* following is setting part
*****************************************/


func (this *AppController) AppSettingSignout() {
  //注销session
  //修改数据库数据表user中相应的isactive字段
  u := this.GetSession("username")
  this.DelSession("username")
  b := models.JsSign{State: 0}
  if models.Crud.User.SetUnActive(u.(string)) == true {
    b.State = 1
  }
  this.Data["json"] = b
  this.ServeJSON()

}

func (this *AppController) AppSettingEdit() {
  b := models.JsUifSign{
    Existnick: 1,
    Existemail: 1,
  }
  var userinfo models.JsUif
  resbody := this.Ctx.Input.RequestBody
  models.App.Setting.Edit(&userinfo, resbody, &b)
  this.Data["json"] = b
  this.ServeJSON()
}

func (this *AppController) AppSettingUpload() {
  b := models.JsSign{State: 0}
  var userinfo models.JsUif
  resbody := this.Ctx.Input.RequestBody
  if models.CwJSON.Unmarshal(resbody, &userinfo) == true {
    if models.Crud.Uif.UpdateIcon(&userinfo) == true {
      b.State = 1
    }
  }
  this.Data["json"] = b
  this.ServeJSON()

}

