package models


func EditExistHelper(userinfo *UserInfoJSON, errmsg *UserinfoErr) {
  if ExistNickname(userinfo.Nickname) == true {
    errmsg.Existnick = 1
  } else {
    errmsg.Existnick = 0
  }

  if ExistEmail(userinfo.Email) == true {
    errmsg.Existemail = 1
  } else {
    errmsg.Existemail = 0
  }

}

func EditHelper(userinfo *UserInfoJSON, resbody []byte, errmsg *UserinfoErr) bool {
  if AnalyzeUserInfoJson(userinfo, resbody) == false {
    return false
  }

  EditExistHelper(userinfo, errmsg)

  if UpdateUserInfo(userinfo) == false {
    return false
  }

  return true
}
