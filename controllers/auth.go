package controllers

import (
	"github.com/astaxie/beego"
  "catwalk/models"
  "html/template"
  "fmt"
)

type AuthController struct {
	beego.Controller
}

func (this *AuthController) AuthCaptcha() {
  idkey, captcha := models.App.Captcha.Create()
  this.SetSession("idkey", idkey)
  this.Data["json"] = template.URL(captcha)
  this.ServeJSON()
}

func (this *AuthController) AuthSignin() {
  b := models.JsSign{State: 0}
  var user models.JsUser
  ok := models.CwJSON.Unmarshal(this.Ctx.Input.RequestBody, &user)
  fmt.Println("user!!!!!!!!!!!!!!!!!!!!!!!!!!: ", user)
  if !ok {
    this.Data["json"] = b
    this.ServeJSON()
    return
  }

  ok = models.Crud.User.Verify(&user) && models.App.Captcha.Verify(this.GetSession("idkey").(string), string(user.Captchainput))

  if !ok {
    this.Data["json"] = b
    this.ServeJSON()
    return
  }

  this.SetSession("username", user.Username)
  b = models.JsSign{State: 1}
  this.Data["json"] = b
  this.ServeJSON()

}

func (this *AuthController) AuthSignup() {
  b := models.JsSign{State: 0}
  resbody := this.Ctx.Input.RequestBody
  idkey := this.GetSession("idkey").(string)
  username := models.App.Sign.Up(&b, resbody, idkey)
  if b.State == 1 {
    this.SetSession("username", username)
  }
  this.Data["json"] = b
  this.ServeJSON()
}
