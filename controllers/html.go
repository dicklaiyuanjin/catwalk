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

func (this *HtmlController) SigninForm() {
  si := signInfo{
    btnsignid: "signinbtn",
    btnclass: "btn-primary",
    btnvalue: "Sign in",
    jsfile: "signin.js",
  }
  signformHelper(this, &si)
}

func (this *HtmlController) SignupForm() {
  si := signInfo{
    btnsignid: "signupbtn",
    btnclass: "btn-success",
    btnvalue: "Sign up",
    jsfile: "signup.js",
  }
  signformHelper(this, &si)

}

func (this *HtmlController) NotFound() {
	this.TplName = "notfound.tpl"
}


func captchaHelper(this *HtmlController) {
  idkey, captcha := models.CwCaptcha.Create()
  this.SetSession("idkey", idkey)
  this.Data["captcha"] = template.URL(captcha)
}

type signInfo struct {
  btnsignid string
  btnclass string
  btnvalue string
  jsfile string
}

func signformHelper(this *HtmlController, si *signInfo) {
  if this.GetSession("username") != nil {
    this.Ctx.Redirect(302, "/app")
  } else {
    this.Data["btnsignid"] = si.btnsignid
    this.Data["btnclass"] = si.btnclass
    this.Data["btnvalue"] = si.btnvalue
    this.Data["jsfile"] = si.jsfile
    captchaHelper(this)
    this.TplName = "sign.tpl"
  }
}
