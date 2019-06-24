package controllers

import (
	"github.com/astaxie/beego"
)

type HtmlController struct {
	beego.Controller
}

func (c *HtmlController) Index() {
	c.TplName = "index.tpl"
}

func (c *HtmlController) SigninForm() {
	c.TplName = "signin.tpl"
}

func (c *HtmlController) SignupForm() {
	c.TplName = "signup.tpl"
}

func (c *HtmlController) App() {
	c.TplName = "app.tpl"
}

func (c *HtmlController) NotFound() {
	c.TplName = "notfound.tpl"
}
