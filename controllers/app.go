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
    //同意之后，将信息放入好友列表
    //除了将信息发送回发送方，还要发送成功信息给对方，这个需要依靠websocket实现
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
    //拒绝之后，不需进行数据库操作，也不需要告诉对方
    b.State = 1
  }
  this.Data["json"] = b
  this.ServeJSON()
}


