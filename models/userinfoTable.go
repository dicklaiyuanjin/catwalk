package models

import (
  "github.com/astaxie/beego/orm"
)

func InsertUserInfo(u *JsUif) bool {
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

func ReadUserInfoUsername(nickname string) string {
  o := orm.NewOrm()
  o.Using("default")

  usr := Userinfo{Nickname: nickname}

  err := o.Read(&usr, "Nickname")

  if err == orm.ErrNoRows || err == orm.ErrMissPK {
    return ""
  } else {
    return usr.Username
  }
}

func ReadUserInfo(u *JsUif, key string) bool{
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

func UpdateUserInfo(u *JsUif) bool {
  o := orm.NewOrm()
  o.Using("default")

  usr := Userinfo{Username: u.Username}

  if o.Read(&usr) == nil {
    usr.Nickname = u.Nickname
    usr.Email = u.Email
    usr.Motto = u.Motto

    if num, err := o.Update(&usr, "Nickname", "Email", "Motto"); err == nil {
      if num == 1 {
        return true
      }
    }
  }

  return false
}

func UpdateIconOfUserInfo(u *JsUif) bool {
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

func ExistNickname(nickname string) bool {
  o := orm.NewOrm()
  o.Using("default")

  usr := Userinfo{Nickname: nickname}

  err := o.Read(&usr, "Nickname")

  if err == orm.ErrNoRows || err == orm.ErrMissPK {
    return false
  } else {
    return true
  }
}

func ExistEmail(email string) bool {
  o := orm.NewOrm()
  o.Using("default")

  usr := Userinfo{Email: email}

  err := o.Read(&usr, "Email")

  if err == orm.ErrNoRows || err == orm.ErrMissPK {
    return false
  } else {
    return true
  }
}
