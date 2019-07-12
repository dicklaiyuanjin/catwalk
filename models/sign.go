package models

func AuthSignupHelper(b *JsSign, resbody []byte, idkey string) string {
  var user *JsUser
  var userinfo JsUif
  if CwJSON.Unmarshal(resbody, user) == false {
    return ""
  }

  if Crud.User.Exist(user.Username) == true {
    return ""
  }

  if CwCaptcha.Verify(idkey, user.Captchainput) == false {
    return ""
  }

  if Crud.User.Insert(user) == false {
    return ""
  }

  userinfo.Username = user.Username

  if InsertUserInfo(&userinfo) == false {
    return ""
  }


  Crud.User.SetActive(user.Username)
  b.State = 1
  return user.Username
}
