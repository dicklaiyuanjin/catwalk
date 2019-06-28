package controllers

import (
  "github.com/astaxie/beego"
  "catwalk/models"
  "html/template"
)

type HtmlController struct {
	beego.Controller
}

func (this *HtmlController) Index() {
	this.TplName = "index.tpl"
}

func captchaHelper(this *HtmlController) {
  idkey, captcha := models.CaptchaCreate()
  this.SetSession("idkey", idkey)
  this.Data["captcha"] = template.URL(captcha)
}

func (this *HtmlController) SigninForm() {
  if this.GetSession("username") != nil {
    this.Ctx.Redirect(302, "/app")
  } else {
    captchaHelper(this)
    this.TplName = "signin.tpl"
  }
}

func (this *HtmlController) SignupForm() {
  if this.GetSession("username") != nil {
    this.Ctx.Redirect(302, "/app")
  } else {
    captchaHelper(this)
    this.TplName = "signup.tpl"
  }
}

func (this *HtmlController) App() {
	this.TplName = "app.tpl"
}

func (this *HtmlController) NotFound() {
	this.TplName = "notfound.tpl"
}
