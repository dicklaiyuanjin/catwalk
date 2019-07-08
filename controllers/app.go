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

  //invitation part
  //作为reciver，获得所有sender发送给自己的invitation
  var inviArr []models.InvitationJSON
  if models.ReadInvitation(&inviArr, userinfo.Nickname, "Receiver") == true {
    fmt.Println(inviArr)
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
  b := models.SignErr{State: 0}
  if models.SetUserUnActive(u.(string)) == true {
    b.State = 1
  }
  this.Data["json"] = b
  this.ServeJSON()

}

func (this *AppController) AppSettingEdit() {
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

func (this *AppController) AppSettingUpload() {
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

/******************************************
* following is invitation part
******************************************/

func (this *AppController) AppInvitationAgree() {
  b := models.SignErr{State: 0}
  var invitation models.InvitationJSON
  resbody := this.Ctx.Input.RequestBody
  if models.AnalyzeInvitationJson(&invitation, resbody) == true {
    fmt.Println("agree!!!!!!!!!!!!!: ", invitation)
    b.State = 1
  }
  this.Data["json"] = b
  this.ServeJSON()
}

func (this *AppController) AppInvitationRefuse() {
  b := models.SignErr{State: 0}
  var invitation models.InvitationJSON
  resbody := this.Ctx.Input.RequestBody
  if models.AnalyzeInvitationJson(&invitation, resbody) == true {
    fmt.Println("refuse!!!!!!!!!!!!!!!:", invitation)
    b.State = 1
  }
  this.Data["json"] = b
  this.ServeJSON()
}

/*
func (this *AppController) AppInvitationSend() {
  b := models.SignErr{State: 0}
  var invitation models.InvitationJSON
  resbody := this.Ctx.Input.RequestBody
  if models.AnalyzeInvitationJson(&invitation, resbody) == true {
    if models.InsertInvitation(&invitation) == true {
      //假设接收者在线，尝试发送给接收者

      //设置成功与否的标志
      b.State = 1
    }
  }
  this.Data["json"] = b
  this.ServeJSON()
}
*/

