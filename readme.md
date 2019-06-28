# Catwalk
* Catwalk is a project to implement an IM system
* For me, the object of creating this project is to study building web application with golang
* As we all known, building web app is not easy. It needs many skills as well as tools.

# why use name catwalk
* I love cats, they are cute and funny, which makes me feel happy.
* The process of study is difficult. I hope I can get happy from it.

# the skills and tools
* [beego](https://github.com/astaxie/beego)


# Router design
| url | ctlr | method |
| :-- | :--- | :----- |
| / | Html | get:Index |
| /signinform | Html | get:SigninForm |
| /signupform | Html | get:SignupForm |
| /app | Html | get:App |
| /notfound | Html | get:NotFound |
| /auth/captcha | Auth | post:AuthCaptcha |
| /auth/signin | Auth | post:AuthSignin |
| /auth/signup | Auth | post:AuthSignup |
| /ws/join/user | Ws | get:JoinUser |

# controller design
* Html: return some html page.
```golang
func (this *HtmlController) Index()
func (this *HtmlController) SigninForm()
func (this *HtmlController) SignupForm()
func (this *HtmlController) App()
func (this *HtmlController) NotFound()

//other thing, help create a captcha data
func captchaHelper(this *HtmlController)
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
* Ws: handle websocket request.
* other thing: helper func will help handle some problem.

# models
## database design
* table:
```sql
CREATE TABLE `user` (
	`uid` INT(10) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(64) NOT NULL UNIQUE,
	`password` BLOB NOT NULL,
	`isactive` INT(1),
	PRIMARY KEY (`uid`)
);
```
* dbmodels: use "github.com/astaxie/beego/orm"
```golang
type User struct {
  Uid int `orm:"pk"`
  Username string `orm:"column(username)"`
  Password string `orm:column(password)`
  Isactive bool `orm:column(isactive)`
}
```
* captcha: use "github.com/mojocn/base64Captcha"
```golang
//return idkey and captcha(base64string)
func CaptchaCreate() (string, string)

func VerifyCaptcha(idkey, verifyValue string) bool
```
* mysqlHandler(this part will be optimized): use follow tools
  ** "golang.org/x/crypto/scrypt"
  ** "github.com/astaxie/beego/orm"
  ** "github.com/go-sql-driver/mysql"
```golang
//username exist or not
func ExistUsername(username string) bool

//insert user info to db
func InsertUser(u *UserJSON) bool

//compare two []byte equal or not
func byteSliceEqual(a, b []byte) bool

func VerifyUser(u *UserJSON) bool

func SetUserActive(u string) bool

func SetUserUnActive(u string) bool
```
* sign models
```golang
//UserJSON is used to save data from front-end
type UserJSON struct {
  Username string `json:"username"`
  Password string `json:"password"`
  Captchainput string `json:"captchainput"`
}

/*
 * state=0: info error
 * state=1: pass
 * this type is used on AuthController.AuthSignin and AuthSignup
 * front-end will check this data, then do sth
 */
type SignErr struct {
  State int `json:"state"`
}

//save data from front-end to UserJSON
func AnalyzeUserJson(user *UserJSON, resbody []byte) bool
```
