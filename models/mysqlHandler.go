package models

import (
  "golang.org/x/crypto/scrypt"
  "github.com/astaxie/beego/orm"
   _ "github.com/go-sql-driver/mysql"
)

func init() {
  orm.RegisterDriver("mysql", orm.DRMySQL)
  orm.RegisterDataBase("default", "mysql", "dick:12345678@/catwalk?charset=utf8")
  orm.Debug = true
}

func ExistUsername(username string) bool {
  o := orm.NewOrm()
  o.Using("default")

  user := User{Username: username}
  err := o.Read(&user)

  if err == orm.ErrNoRows || err == orm.ErrMissPK {
    return false
  } else {
    return true
  }
}

func InsertUser(u *UserJSON) bool {
  if ExistUsername(u.Username) {
    return false
  }

  o := orm.NewOrm()
  o.Using("default")
  var user User
  user.Username = u.Username
  pwd, err := scrypt.Key([]byte(u.Password), Salt, 16384, 8, 1, 32)
  if err != nil {
    return false
  }
  user.Password = string(pwd)
  user.Isactive = false

  _, err = o.Insert(&user)
  if err == nil {
    return true
  }

  return false
}

func byteSliceEqual(a, b []byte) bool {
    if len(a) != len(b) {
        return false
    }

    if (a == nil) != (b == nil) {
        return false
    }

    for i, v := range a {
        if v != b[i] {
            return false
        }
    }

    return true
}

func VerifyUser(u *UserJSON) bool {
  o := orm.NewOrm()
  o.Using("default")

  user := User{Username: u.Username}
  err := o.Read(&user, "Username")

  if err == orm.ErrNoRows || err == orm.ErrMissPK {
    return false
  } else {
    if user.Isactive {
      return false
    }

    pwd, err := scrypt.Key([]byte(u.Password), Salt, 16384, 8, 1, 32)
    if err != nil {
      return false
    }

    return byteSliceEqual(pwd, []byte(user.Password))
  }
}

func SetUserActive(u string) bool {
  o := orm.NewOrm()
  o.Using("default")

  user := User{Username: u}
  err := o.Read(&user, "Username")

  if err == orm.ErrNoRows || err == orm.ErrMissPK {
    return false
  } else {
    user.Isactive = true
    if _, err = o.Update(&user); err == nil {
      return true
    } else {
      return false
    }
  }
}

func SetUserUnActive(u string) bool {
  o := orm.NewOrm()
  o.Using("default")

  user := User{Username: u}
  err := o.Read(&user, "Username")

  if err == orm.ErrNoRows || err == orm.ErrMissPK {
    return false
  } else {
    user.Isactive = false
    if _, err = o.Update(&user); err == nil {
      return true
    } else {
      return false
    }
  }
}
