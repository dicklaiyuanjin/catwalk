package models

func AuthSignupHelper(b *JsSign, resbody []byte, idkey string) string {
  var user *JsUser
  var userinfo JsUif
  if CwJSON.Unmarshal(resbody, user) == false {
    return ""
  }

  if ExistUsername(user.Username) == true {
    return ""
  }

  if CwCaptcha.Verify(idkey, user.Captchainput) == false {
    return ""
  }

  if InsertUser(user) == false {
    return ""
  }

  userinfo.Username = user.Username

  if InsertUserInfo(&userinfo) == false {
    return ""
  }


  SetUserActive(user.Username)
  b.State = 1
  return user.Username
}
