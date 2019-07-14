package models

import (
  "github.com/mojocn/base64Captcha"
)

type AppModel struct {
  Sign *signForm
  Setting *setting
  Captcha *captchaTool
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

  if App.Captcha.Verify(idkey, string(user.Captchainput)) == false {
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

/*********************************************************
 * captchaTool
 ********************************************************/
type captchaTool struct {
  name string
}

func (c *captchaTool) Create() (string, string) {
  var config = base64Captcha.ConfigCharacter {
    Height:             60,
    Width:              240,
    Mode:               base64Captcha.CaptchaModeNumber,
    ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
    ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
    IsShowHollowLine:   true,
    IsShowNoiseDot:     true,
    IsShowNoiseText:    true,
    IsShowSlimeLine:    true,
    IsShowSineLine:     true,
    CaptchaLen:         6,
  }

  idKey, captcha := base64Captcha.GenerateCaptcha("", config)
  base64string := base64Captcha.CaptchaWriteToBase64Encoding(captcha)
  return idKey, base64string
}

func (c *captchaTool) Verify(idkey, verifyValue string) bool {
  return base64Captcha.VerifyCaptcha(idkey, verifyValue)
}
