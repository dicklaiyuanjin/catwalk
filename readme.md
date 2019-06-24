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
| /auth/signin | Auth | post:AuthSignin |
| /auth/signup | Auth | post:AuthSignup |
| /ws/join/user | Ws | get:JoinUser |

# controller design
* Html: return some html page.
* Auth: do something about authorization, such as authorize user signin.
* Ws: handle websocket request.
