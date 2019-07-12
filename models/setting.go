package models


func EditExistHelper(userinfo *JsUif, errmsg *JsUifSign) {
  if Crud.Uif.ExistNickname(userinfo.Nickname) == true {
    errmsg.Existnick = 1
  } else {
    errmsg.Existnick = 0
  }

  if Crud.Uif.ExistEmail(userinfo.Email) == true {
    errmsg.Existemail = 1
  } else {
    errmsg.Existemail = 0
  }

}

func EditHelper(userinfo *JsUif, resbody []byte, errmsg *JsUifSign) bool {
  if CwJSON.Unmarshal(resbody, userinfo) == false {
    return false
  }

  EditExistHelper(userinfo, errmsg)

  if Crud.Uif.Update(userinfo) == false {
    return false
  }

  return true
}
