package controllers

import (
	"github.com/astaxie/beego"
  "catwalk/models"
)

type AppController struct {
	beego.Controller
}

func rToI(this *AppController) {
  this.Ctx.Redirect(302, "/")
}

func (this *AppController) App() {
  usr := this.GetSession("username").(string)
  if usr == "" {
    rToI(this)
    return
  }

  ok := models.Crud.User.SetActive(usr);
  if !ok {
    rToI(this)
    return
  }

  //setting part
  var userinfo models.JsUif
  userinfo.Username = usr
  ok = models.Crud.Uif.Read(&userinfo, "username")
  if !ok {
    rToI(this)
    return
  }
  this.Data["userinfo"] = userinfo


  //invitation part
  //作为reciver，获得所有sender发送给自己的invitation
  var inviArr []models.JsIvtt
  ok = models.Crud.Invitation.ReadList(&inviArr, userinfo.Username, "Receiver")
  if !ok {
    rToI(this)
    return
  }
  this.Data["Invitations"] = inviArr


  //chatroom part
  //friboxs section
  //先获取friednlist，根据好友用户名查其详细信息
  //获取friendinfo列表，发送到chatroom中显示
  var frilist []models.JsFl
  ok = models.Crud.FriendList.ReadList(&frilist, userinfo.Username)
  if !ok {
    rToI(this)
    return
  }

  var fiflist []models.JsUif
  ok = models.Crud.Uif.ReadFifList(&fiflist, &frilist)
  if !ok {
    rToI(this)
    return
  }

  this.Data["Fiflist"] = fiflist

  this.TplName = "app.tpl"
}

/*****************************************
* following is setting part
*****************************************/


func (this *AppController) AppSettingSignout() {
  //注销session
  //修改数据库数据表user中相应的isactive字段
  u := this.GetSession("username")
  this.DestroySession()
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

