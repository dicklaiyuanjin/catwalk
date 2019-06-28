package models

import (
  "encoding/json"
)

type UserJSON struct {
  Username string `json:"username"`
  Password string `json:"password"`
  Captchainput string `json:"captchainput"`
}



func AnalyzeUserJson(user *UserJSON, resbody []byte) bool {
  if err := json.Unmarshal(resbody, user); err == nil {
    return true
  }
  return false
}

/*
 * state=0: info error 
 * state=1: pass
 */
type SignErr struct {
  State int `json:"state"`
}



