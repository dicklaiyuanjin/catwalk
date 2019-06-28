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
  idkey, captcha := models.CaptchaCreate()
  this.SetSession("idkey", idkey)
  this.Data["json"] = template.URL(captcha)
  this.ServeJSON()
}

func (this *AuthController) AuthSignin() {
  b := models.SignErr{State: 0}
  var user models.UserJSON
  if models.AnalyzeUserJson(&user, this.Ctx.Input.RequestBody) == true {
    if models.VerifyUser(&user) == true &&
        models.VerifyCaptcha(this.GetSession("idkey").(string), user.Captchainput) {
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
  var user models.UserJSON
  if models.AnalyzeUserJson(&user, this.Ctx.Input.RequestBody) == true {
    if models.VerifyCaptcha(this.GetSession("idkey").(string), user.Captchainput) {
      if models.ExistUsername(user.Username) == false {
        if models.InsertUser(&user) == true {
          models.SetUserActive(user.Username)
          this.SetSession("username", user.Username)
          b = models.SignErr{State: 1}
        }
      }
    }
  }
  this.Data["json"] = b
  this.ServeJSON()
}
