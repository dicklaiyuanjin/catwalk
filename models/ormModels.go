package models

import (
  "github.com/astaxie/beego/orm"
)

type User struct {
  Uid int `orm:"pk"`
  Username string `orm:"column(username)"`
  Password string `orm:column(password)`
  Isactive int `orm:column(isactive)`
}

type Userinfo struct {
  Username string `orm:"pk"`
  Nickname string `orm:"column(nickname)"`
  Email string `orm:"column(email)"`
  Motto string `orm:"column(motto)"`
  Icon string `orm:"column(icon)"`
}

var Salt []byte

func init() {
  Salt = []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}
  orm.RegisterModel(new(User), new(Userinfo))
}
