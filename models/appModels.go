package models

type AppModel struct {
  Sign *signForm
  Setting *setting
}

var App AppModel

/**********************************************
 * signform
 *********************************************/
type signForm struct {
  name string
}

func (s *signForm) Up(b *JsSign, resbody []byte, idkey string) string {
  var user JsUser
  var userinfo JsUif
  if CwJSON.Unmarshal(resbody, &user) == false {
    return ""
  }

  if Crud.User.Exist(user.Username) == true {
    return ""
  }

  if CwCaptcha.Verify(idkey, user.Captchainput) == false {
    return ""
  }

  if Crud.User.Insert(&user) == false {
    return ""
  }

  userinfo.Username = user.Username
  userinfo.Nickname = user.Username

  if Crud.Uif.Insert(&userinfo) == false {
    return ""
  }


  Crud.User.SetActive(user.Username)
  b.State = 1
  return user.Username
}


/***************************************************
 * setting
 **************************************************/
type setting struct {
  name string
}

func (s *setting) SetExistSign(userinfo *JsUif, sign *JsUifSign) {
  if Crud.Uif.Exist(userinfo.Nickname, "Nickname") == true {
    sign.Existnick = 1
  } else {
    sign.Existnick = 0
  }

  if Crud.Uif.Exist(userinfo.Email, "Email") == true {
    sign.Existemail = 1
  } else {
    sign.Existemail = 0
  }
}

func (s *setting) Edit(userinfo *JsUif, resbody []byte, sign *JsUifSign) bool {
  if CwJSON.Unmarshal(resbody, userinfo) == false {
    return false
  }

  s.SetExistSign(userinfo, sign)

  if Crud.Uif.Update(userinfo) == false {
    return false
  }

  return true
}
