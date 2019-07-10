package models

func AuthSignupHelper(b *SignErr, resbody []byte, idkey string) string {
  var user UserJSON
  var userinfo UserInfoJSON
  if AnalyzeUserJson(&user, resbody) == false {
    return ""
  }

  if ExistUsername(user.Username) == true {
    return ""
  }

  if CwCaptcha.Verify(idkey, user.Captchainput) == false {
    return ""
  }

  if InsertUser(&user) == false {
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
