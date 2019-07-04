# Catwalk
* Catwalk is a project to implement an IM system (including User Authentication)
* For me, the object of creating this project is to study building web application with golang
* As we all known, building web app is not easy. It needs many skills as well as tools.

# why use name catwalk
* I love cats, they are cute and funny, which makes me feel happy.
* The process of study is difficult. I hope I can get happy from it.

# the skills and tools
* [beego](https://github.com/astaxie/beego)

# files:
```
├── conf
│   └── app.conf
├── controllers
│   ├── app.go
│   ├── auth.go
│   ├── html.go
│   └── ws.go
├── main.go
├── models
│   ├── captcha.go
│   ├── db_mysql.go
│   ├── JSONModels.go
│   ├── ormModels.go
│   ├── sign.go
│   ├── userinfoTable.go
│   └── userTable.go
├── readme.md
├── routers
│   └── router.go
├── static
│   ├── css
│   │   ├── basic.css
│   │   ├── setting.css
│   │   └── sign.css
│   ├── img
│   │   └── icon.png
│   └── js
│       ├── app
│       │   ├── basic.js
│       │   ├── setting.js
│       │   └── start.js
│       ├── captcha.js
│       ├── signin.js
│       ├── sign.js
│       └── signup.js
├── tests
│   └── default_test.go
└── views
    ├── app.tpl
    ├── index.tpl
    ├── notfound.tpl
    ├── sign.tpl
    └── tpl
        ├── app
        │   ├── chatroom.tpl
        │   ├── head.tpl
        │   ├── invitation.tpl
        │   ├── setting.tpl
        │   └── tail.tpl
        ├── head.tpl
        ├── signform.tpl
        └── tail.tpl
```

# back-end

## Router design
| url | ctlr | method |
| :-- | :--- | :----- |
| / | Html | get:Index |
| /signinform | Html | get:SigninForm |
| /signupform | Html | get:SignupForm |
| /notfound | Html | get:NotFound |
| /auth/captcha | Auth | post:AuthCaptcha |
| /auth/signin | Auth | post:AuthSignin |
| /auth/signup | Auth | post:AuthSignup |
| /app | App | get:App |
| /app/signout | App | get:AppSignout |
| /app/edit | App | post:AppEdit |
| /app/upload | App | post:AppUpload |
| /ws/join/user | Ws | get:JoinUser |

## controller design
* Html: return some html page.
```golang
func (this *HtmlController) Index()
func (this *HtmlController) SigninForm()
func (this *HtmlController) SignupForm()
func (this *HtmlController) NotFound()

//other thing, help create a captcha data
func captchaHelper(this *HtmlController)
func signformHelper(this *HtmlController, si *signInfo)
```
* Auth: do something about authorization, such as authorize user signin.
```golang
//when user clicks picture of captcha, this func will refresh the picture.
func (this *AuthController) AuthCaptcha()

/*
 * this func will be optimized, the processes are follow:
 * Analyze userinfo from front-end.
 * Verify userinfo and captcha
 * if userinfo all right: set session, active the user, return success msg to front-end
 * if not:return fail msg to front-end
 */
func (this *AuthController) AuthSignin()


/*
 * this func will be optimized, the processes are follow:
 * Analyze userinfo from front-end
 * Verify captcha
 * check username exist or not
 * if userinfo all right: add user info to database, active the user, set session, return success msg to front-end
 * if not: return fail msg to front-end
 */
func (this *AuthController) AuthSignup()
```
* App: do something about app
```golang
//get user info and put then in app page,then return to user
func (this *AppController) App()

//when user sign out, this controller will delete the user session
//then set user inactive on database
func (this *AppController) AppSignout()

//when user edit his/her user infomation and submit
//this controller will modify the relative data on database
func (this *AppController) AppEdit()

//when user submit new user icon, this controller will update user infomation relatively
func (this *AppController) AppUpload()
```
* Ws: handle websocket request.
* other thing: helper func will help handle some problem.

## models
* database table:
```sql
CREATE TABLE `user` (
	`uid` INT(10) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(64) NOT NULL UNIQUE,
	`password` BLOB NOT NULL,
	`isactive` INT(1),
	PRIMARY KEY (`uid`)
);

CREATE TABLE `userinfo` (
	`username` VARCHAR(64) NOT NULL,
	`nickname` VARCHAR(64),
	`email` VARCHAR(255),
	`motto` VARCHAR(255),
	`icon` BLOB,
	PRIMARY KEY(`username`)
);
```
* models:JSONModels.go (This part defines the structure to handle the data between front-end and back-end. Besides, it also has function to handle json data)
```golang
type UserJSON struct {
  Username string `json:"username"`
  Password string `json:"password"`
  Captchainput string `json:"captchainput"`
}

/*
 * state=0: info error
 * state=1: pass
 */
type SignErr struct {
  State int `json:"state"`
}

type UserInfoJSON struct {
  Username string `json:"username"`
  Nickname string `json:"nickname"`
  Email string `json:"email"`
  Motto string `json"motto"`
  Icon string `json"icon"`
}

func AnalyzeUserJson(user *UserJSON, resbody []byte) bool

func AnalyzeUserInfoJson(userinfo *UserInfoJSON, resbody []byte) bool
```
* models: ormModels (This part defines then structure to reflect database table, also some orm handler)
```golang
type User struct {
  Uid int `orm:"pk"`
  Username string `orm:"column(username)"`
  Password string `orm:column(password)`
  Isactive int `orm:column(isactive)`
}

type Userinfo struct {
  Username string `orm:"pk"`
  Nickname string `orm:"column(nickname)"`
  Email string `orm:"column(email)"`
  Motto string `orm:"column(motto)"`
  Icon string `orm:"column(icon)"`
}
```
* models: db_mysql.go (this part register the mysql driver and orm driver)
* models: captcha.go (use github.com/mojocn/base64Captcha)
```golang
//return idkey and captcha encoding with base64
func CaptchaCreate() (string, string)

func VerifyCaptcha(idkey, verifyValue string) bool
```
* models: sign.go
```golang
//help User Authentication
func AuthSignupHelper(b *SignErr, resbody []byte, idkey string) string
```
* models: userTable.go (this part defines "crud" of table user)
```golang
func ExistUsername(username string) bool

func InsertUser(u *UserJSON) bool

func VerifyUser(u *UserJSON) bool

func SetUserActive(u string) bool

func SetUserUnActive(u string) bool
```
* models: userinfoTable.go (this part defines "crud" of table userinfo)
```golang
func InsertUserInfo(u *UserInfoJSON) bool

func ReadUserInfo(u *UserInfoJSON, key string) bool

func UpdateUserInfo(u *UserInfoJSON) bool

func UpdateIconOfUserInfo(u *UserInfoJSON) bool
```

# front-end
* use tools:  bootstrap 3.3.7
* views design:
```
├── app.tpl
├── index.tpl
├── notfound.tpl
├── sign.tpl
└── tpl
    ├── app
    │   ├── chatroom.tpl
    │   ├── head.tpl
    │   ├── invitation.tpl
    │   ├── setting.tpl
    │   └── tail.tpl
    ├── head.tpl
    ├── signform.tpl
    └── tail.tpl
```

* static resouces:
```
├── css
│   ├── basic.css
│   ├── setting.css
│   └── sign.css
├── img
│   └── icon.png
└── js
    ├── app
    │   ├── basic.js
    │   ├── setting.js
    ├── captcha.js
    ├── signin.js
    ├── sign.js
    └── signup.js
```
* about the js file:
1. captcha.js: refresh the captcha picture
2. sign.js: the function package to help signin or signup
3. signin.js: use sign.js to handle signin request
4. signup.js : use sign.js to handle signup request
5. app/basic.js: some function app will use
6. app/setting.js: handle the action from user in setting page

