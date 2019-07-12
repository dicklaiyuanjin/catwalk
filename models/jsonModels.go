package models

import (
  "encoding/json"
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
  Captchainput string `json:"captchainput"`
}


/*************************************
 * JsUif(Uif: userinfo)
 ************************************/
type JsUif struct {
  Username string `json:"username"`
  Nickname string `json:"nickname"`
  Email string `json:"email"`
  Motto string `json"motto"`
  Icon string `json"icon"`
}




/********************************************
 * JsIvtt
 *******************************************/
//Sender and Receiver should be username
type JsIvtt struct {
  Sender string `json:"sender"`
  Receiver string `json:"receiver"`
  Msg string `json:"msg"`
}



/********************************************
 * JsFl
 *******************************************/
type JsFl struct {
  Username string `json:"username"`
  Friusername string `json:"friusername"`
}


