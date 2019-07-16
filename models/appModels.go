package models

import (
  "github.com/mojocn/base64Captcha"
)

type AppModel struct {
  Sign *signForm
  Setting *setting
  Captcha *captchaTool
  Msg *msgTool
  Rpl *rplTool
  Del *delTool
}

var App AppModel
/********************************************
 * del: delete friend tool
 *******************************************/
type delTool struct {
  name string
}

func (dt *delTool) CheckAndDel(d *JsDel) bool {
  ok := (Crud.User.Exist(d.Sender) && Crud.User.Exist(d.Exfri))
  if !ok { return false }

  fl := JsFl{
    Username: d.Sender,
    Friusername: d.Exfri,
  }

  ok = Crud.FriendList.ExistList(&fl)
  if !ok { return false }

  ok = Crud.FriendList.Delete(&fl)
  if !ok { return false }

  return true
}



/***********************************************
 * rpl:reply tool
 **********************************************/
type rplTool struct {
  name string
}

func (rt *rplTool) Check(r *JsRpl) bool {
  ok := Crud.Invitation.Exist(r.Obj, r.Me)
  if !ok { return false }

  ok = Crud.Invitation.Delete(r.Obj, r.Me)
  if !ok { return false }

  Crud.Invitation.Delete(r.Me, r.Obj)

  return true
}

func (rt *rplTool) AddFri(r *JsRpl) bool {
  ok := Crud.FriendList.Insert(&JsFl{
    Username: r.Me,
    Friusername: r.Obj,
  })
  if !ok { return false }
  return true
}

func (rt *rplTool) Marshal(src string) ([]byte, bool) {
  Src := JsUif{Username: src}
  ok := Crud.Uif.Read(&Src, "Username")
  if !ok { return nil, false }

  ws := WsData{
    Code: 2,
    Fif: Src,
  }

  data, ok := CwJSON.Marshal(ws)
  return data, ok
}

/***********************************************
 * msg tool
 **********************************************/
type msgTool struct {
  name string
}

func (mt *msgTool) CheckAndInsert(m *JsMsg) bool {
  ok := Crud.User.Exist(m.Sender) && Crud.User.Exist(m.Receiver)
  if !ok { return false }

  ok = Crud.FriendList.ExistList(&JsFl{
    Username: m.Sender,
    Friusername: m.Receiver,
  })
  if !ok { return false }

  ok = (m.Content != "")
  if !ok { return false }

  ok = Crud.Msg.Insert(m)
  if !ok { return false }

  return true
}



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
  userinfo.Icon = "/static/img/icon.png"

  if Crud.Uif.Insert(&userinfo) == false {
    return ""
  }


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
