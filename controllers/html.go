package controllers

import (
	"github.com/astaxie/beego"
)

type HtmlController struct {
	beego.Controller
}

func (c *HtmlController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *HtmlController) SigninForm() {

}

func (c *HtmlController) SignupForm() {

}

func (c *HtmlController) App() {

}
