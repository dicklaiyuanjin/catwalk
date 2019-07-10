package controllers

import (
	"github.com/astaxie/beego"
  "catwalk/models"
  "html/template"
)

type AuthController struct {
	beego.Controller
}

func (this *AuthController) AuthCaptcha() {
  idkey, captcha := models.CwCaptcha.Create()
  this.SetSession("idkey", idkey)
  this.Data["json"] = template.URL(captcha)
  this.ServeJSON()
}

func (this *AuthController) AuthSignin() {
  b := models.SignErr{State: 0}
  var user models.UserJSON
  if models.AnalyzeUserJson(&user, this.Ctx.Input.RequestBody) == true {
    if models.VerifyUser(&user) == true &&
        models.CwCaptcha.Verify(this.GetSession("idkey").(string), user.Captchainput) {
      this.SetSession("username", user.Username)
      models.SetUserActive(user.Username)
      b = models.SignErr{State: 1}
    }
  }

  this.Data["json"] = b
  this.ServeJSON()

}

func (this *AuthController) AuthSignup() {
  b := models.SignErr{State: 0}
  resbody := this.Ctx.Input.RequestBody
  idkey := this.GetSession("idkey").(string)
  username := models.AuthSignupHelper(&b, resbody, idkey)
  if b.State == 1 {
    this.SetSession("username", username)
  }
  this.Data["json"] = b
  this.ServeJSON()
}
