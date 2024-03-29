package models

import (
  "encoding/json"
  "html/template"
)

type CwJSONModel struct {
  name string
}

var CwJSON *CwJSONModel

func (cj *CwJSONModel) Unmarshal(resbody []byte, v interface{}) bool {
  if err := json.Unmarshal(resbody, v); err == nil {
    return true
  }
  return false
}

func (cj *CwJSONModel) Marshal (v interface{}) ([]byte, bool) {
  if b, err := json.Marshal(v); err == nil {
    return b, true
  }

  return nil, false
}

/*********************************************************
 * the websocket data between back-end and front-end
 * code represent the type
 * 0: Ivtt(invitation)
 * 1: Rpl(reply)
 * 2: fif(friendinfo)
 * 3: Msg(message)
 * 4: Del(delete friend)
 ********************************************************/
type WsData struct {
  Code int `json:"code"`
  Ivtt JsIvtt `json:"ivtt"`
  Rpl JsRpl `json:"rpl"`
  Fif JsUif `json:"fif"`
  Msg JsMsg `json:"msg"`
  Del JsDel `json:"del"`
}


/************************************
 * JsSign
 * state=0: info error
 * state=1: pass
 ************************************/
type JsSign struct {
  State int `json:"state"`
}


/*********************************
 * JsUifSign(Uif: userinfo)
 * value=1 : exist
 * value=0 : not exist
 ********************************/
type JsUifSign struct {
  Existnick int `json:"existnick"`
  Existemail int `json:"existemail"`
}



/*******************************
 * JsUser
 ******************************/
type JsUser struct {
  Username string `json:"username"`
  Password string `json:"password"`
  Captchainput template.URL `json:"captchainput"`
}


/*************************************
 * JsUif(Uif: userinfo)
 ************************************/
type JsUif struct {
  Username string `json:"username"`
  Nickname string `json:"nickname"`
  Email string `json:"email"`
  Motto string `json:"motto"`
  Icon template.URL `json:"icon"`
}




/********************************************
 * JsIvtt(Ivtt:invitation)
 *******************************************/
//Sender and Receiver should be username
type JsIvtt struct {
  Sender string `json:"sender"`
  Receiver string `json:"receiver"`
  Msg string `json:"msg"`
}

/*************************************************
 * JsRpl(Rpl:reply)
 ************************************************/
type JsRpl struct {
  Me string `json:"me"`
  Obj string `json:"obj"`
  Attitude string `json:"attitude"`
}


/********************************************
 * JsFl(Fl:friendlist)
 *******************************************/
type JsFl struct {
  Username string `json:"username"`
  Friusername string `json:"friusername"`
}

/*********************************************
 * JsMsg(Msg: message)
 ********************************************/
type JsMsg struct {
  Sender string `json:"sender"`
  Receiver string `json:"receiver"`
  Content string `json:"content"`
  Sendtime string `json:"sendtime"`
}

/*******************************************
 * JsDel(Del: delete friend)
 ******************************************/
type JsDel struct {
  Sender string `json:"sender"`
  Exfri string `json:"exfri"`
}

/*******************************************
 * JsUifWithFriMsg
 ******************************************/
type JsUifWithFriMsg struct{
  Username string
  Friusername string
  Frinickname string
  Friemail string
  Frimotto string
  Friicon template.URL
  Msg []JsMsg
}





