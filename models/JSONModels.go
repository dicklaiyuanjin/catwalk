package models

import (
  "encoding/json"
)

type UserJSON struct {
  Username string `json:"username"`
  Password string `json:"password"`
  Captchainput string `json:"captchainput"`
}

/*
 * state=0: info error
 * state=1: pass
 */
type SignErr struct {
  State int `json:"state"`
}

type UserInfoJSON struct {
  Username string `json:"username"`
  Nickname string `json:"nickname"`
  Email string `json:"email"`
  Motto string `json"motto"`
  Icon string `json"icon"`
}


func AnalyzeUserJson(user *UserJSON, resbody []byte) bool {
  if err := json.Unmarshal(resbody, user); err == nil {
    return true
  }
  return false
}

func AnalyzeUserInfoJson(userinfo *UserInfoJSON, resbody []byte) bool {
  if err := json.Unmarshal(resbody, userinfo); err == nil {
    return true
  }
  return false
}

