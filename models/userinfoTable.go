package models

import (
  "github.com/astaxie/beego/orm"
)

func InsertUserInfo(u *UserInfoJSON) bool {
  if ExistUsername(u.Username) {
    return false
  }

  o := orm.NewOrm()
  o.Using("default")
  userinfo :=  Userinfo{
    Username: u.Username,
    Nickname: u.Nickname,
    Email: u.Email,
    Motto: u.Motto,
    Icon: u.Icon,
  }

  _, err := o.Insert(&userinfo)
  if err == nil {
    return true
  }

  return false
}

func ReadUserInfo(u *UserInfoJSON, key string) bool{
  o := orm.NewOrm()
  o.Using("default")

  usr := Userinfo{
    Username: u.Username,
    Nickname: u.Nickname,
    Email: u.Email,
    Motto: u.Motto,
    Icon: u.Icon,
  }

  err := o.Read(&usr, key)

  if err == orm.ErrNoRows || err == orm.ErrMissPK {
    return false
  } else {
    u.Username = usr.Username
    u.Nickname = usr.Nickname
    u.Email = usr.Email
    u.Motto = usr.Motto
    u.Icon = usr.Icon
    return true
  }
}

func UpdateUserInfo(u *UserInfoJSON) bool {
  o := orm.NewOrm()
  o.Using("default")

  usr := Userinfo{Username: u.Username}

  if o.Read(&usr) == nil {
    usr.Nickname = u.Nickname
    usr.Email = u.Email
    usr.Motto = u.Motto

    if _, err := o.Update(&usr, "Nickname", "Email", "Motto"); err == nil {
      return true
    }
  }

  return false
}

func UpdateIconOfUserInfo(u *UserInfoJSON) bool {
  o := orm.NewOrm()
  o.Using("default")

  usr := Userinfo{Username: u.Username}

  if o.Read(&usr) == nil {
    usr.Icon = u.Icon
    if _, err := o.Update(&usr, "Icon"); err == nil {
      return true
    }
  }
  return false
}
