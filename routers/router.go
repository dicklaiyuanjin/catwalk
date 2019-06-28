package routers

import (
	"catwalk/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//Html Router: return html page
	beego.Router("/", &controllers.HtmlController{}, "get:Index")
	beego.Router("signinform", &controllers.HtmlController{}, "get:SigninForm")
	beego.Router("signupform", &controllers.HtmlController{}, "get:SignupForm")
	beego.Router("/app", &controllers.HtmlController{}, "get:App")
	beego.Router("/notfound", &controllers.HtmlController{}, "get:NotFound")

	//Auth Router: auth receive info
  beego.Router("/auth/captcha", &controllers.AuthController{}, "post:AuthCaptcha")
	beego.Router("/auth/signin", &controllers.AuthController{}, "post:AuthSignin")
	beego.Router("/auth/signup", &controllers.AuthController{}, "post:AuthSignup")

	//Webocket Router: handle websocket
	beego.Router("/ws/join/user", &controllers.WsController{}, "get:JoinUser")
}
