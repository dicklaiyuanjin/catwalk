package models

import (
  "github.com/astaxie/beego/orm"
  _ "github.com/go-sql-driver/mysql"
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

//Sender and Receiver should be username
type Invitation struct {
  Iid int `orm:"pk"`
  Sender string `orm:"column(sender)"`
  Receiver string `orm:"column(receiver)"`
  Msg string `orm:"column(msg)"`
}

type Friendlist struct {
  Fid int `orm:"pk"`
  Username string `orm:"column(username)"`
  Friusername string `orm:"column(friusername)"`
}

var Salt []byte

func init() {
  Salt = []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}
  orm.RegisterDriver("mysql", orm.DRMySQL)
  orm.RegisterDataBase("default", "mysql", "dick:12345678@/catwalk?charset=utf8")
  orm.Debug = true
  orm.RegisterModel(new(User), new(Userinfo), new(Invitation), new(Friendlist))
}
