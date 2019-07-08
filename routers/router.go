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
	beego.Router("/notfound", &controllers.HtmlController{}, "get:NotFound")

	//Auth Router: auth receive info
  beego.Router("/auth/captcha", &controllers.AuthController{}, "post:AuthCaptcha")
	beego.Router("/auth/signin", &controllers.AuthController{}, "post:AuthSignin")
	beego.Router("/auth/signup", &controllers.AuthController{}, "post:AuthSignup")

  //App Router: handle app and app ajax
	beego.Router("/app", &controllers.AppController{}, "get:App")
  beego.Router("/app/setting/signout", &controllers.AppController{}, "get:AppSettingSignout")
  beego.Router("/app/setting/edit", &controllers.AppController{}, "post:AppSettingEdit")
  beego.Router("/app/setting/upload", &controllers.AppController{}, "post:AppSettingUpload")
  beego.Router("/app/invitation/agree", &controllers.AppController{}, "post:AppInvitationAgree")
  beego.Router("/app/invitation/refuse", &controllers.AppController{}, "post:AppInvitationRefuse")

  //Webocket Router: handle websocket
  beego.Router("/ws/ivtt/join", &controllers.WsController{}, "get:JoinIvttUser")
}
