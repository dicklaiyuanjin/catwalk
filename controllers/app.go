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
  if models.ReadUserInfo(&userinfo, "username") == true {
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
  ok := models.Crud.Invitation.ReadList(inviArr, userinfo.Username, "Receiver")
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
  models.EditHelper(&userinfo, resbody, &b)
  this.Data["json"] = b
  this.ServeJSON()
}

func (this *AppController) AppSettingUpload() {
  b := models.JsSign{State: 0}
  var userinfo models.JsUif
  resbody := this.Ctx.Input.RequestBody
  if models.CwJSON.Unmarshal(resbody, &userinfo) == true {
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
  b := models.JsSign{State: 0}
  var ivtt models.JsIvtt
  resbody := this.Ctx.Input.RequestBody
  if models.CwJSON.Unmarshal(resbody, &ivtt) == true {
    fmt.Println("agree!!!!!!!!!!!!!: ", ivtt)
    //同意之后，将信息放入好友列表
    //除了将信息发送回发送方，还要发送成功信息给对方，这个需要依靠websocket实现
    b.State = 1
  }
  this.Data["json"] = b
  this.ServeJSON()
}

func (this *AppController) AppInvitationRefuse() {
  b := models.JsSign{State: 0}
  var ivtt models.JsIvtt
  resbody := this.Ctx.Input.RequestBody
  if models.CwJSON.Unmarshal(resbody, &ivtt) == true {
    fmt.Println("refuse!!!!!!!!!!!!!!!:", ivtt)
    //拒绝之后，不需进行数据库操作，也不需要告诉对方
    b.State = 1
  }
  this.Data["json"] = b
  this.ServeJSON()
}


