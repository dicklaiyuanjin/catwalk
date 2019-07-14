package models

import (
  "github.com/astaxie/beego/orm"
  "golang.org/x/crypto/scrypt"
)

type CrudModel struct {
  FriendList *flTbl
  Invitation *ivttTbl
  User *userTbl
  Uif *userinfoTbl
}

var Crud CrudModel

/************************************************************************
 * friendlist table
 ***********************************************************************/
type flTbl struct {
  name string
}

//检查是否对称存在
func (fl *flTbl) ExistList(f *JsFl) bool {
  o := orm.NewOrm()
  o.Using("default")

  var fri Friendlist
  num1, err1 := o.QueryTable("friendlist").Filter("username", f.Username).Filter("friusername", f.Friusername).All(&fri)
  num2, err2 := o.QueryTable("friendlist").Filter("username", f.Friusername).Filter("friusername", f.Username).All(&fri)
 if err1 == nil && err2 == nil{
    if num1 == 1 && num2 == 1 {
      return true
    }
  }
  return false
}

//对称插入好友列表（这是一个冗余操作）
func (fl *flTbl) Insert(f *JsFl) bool {
  o := orm.NewOrm()
  o.Using("default")

  fri1 := Friendlist{
    Username: f.Username,
    Friusername: f.Friusername,
  }

  fri2 := Friendlist{
    Username: f.Friusername,
    Friusername: f.Username,
  }

  if fl.ExistList(f) == false && f.Username != f.Friusername {
    _, err1 := o.Insert(&fri1)
    _, err2 := o.Insert(&fri2)
    if err1 == nil && err2 == nil {
      return true
    }
  }

  return false
}

//以username为主体，获取他/她的好友列表
func (fl *flTbl) ReadList(f *[]JsFl, username string) bool {
  o := orm.NewOrm()
  o.Using("default")

  var fris []Friendlist

  _, err := o.QueryTable("friendlist").Filter("username", username).All(&fris)
  if err == nil {
    for _, v := range fris {
      var item JsFl
      item.Username = v.Username
      item.Friusername = v.Friusername
      *f = append(*f, item)
    }
    return true
  }

  return false

}

func (fl *flTbl) ReadId(f *JsFl) (int, bool) {
  o := orm.NewOrm()
  o.Using("default")

  var fris []Friendlist
  num, err := o.QueryTable("friendlist").Filter("username", f.Username).Filter("friusername", f.Friusername).All(&fris)
  if err == nil {
    if num == 1 {
      return fris[0].Fid, true
    }
  }

  return -1, false
}

//对称删除
func (fl *flTbl) Delete(f *JsFl) bool {
  o := orm.NewOrm()
  o.Using("default")

  fid1, ok1 := fl.ReadId(f)
  fid2, ok2 := fl.ReadId(&JsFl{
    Username: f.Username,
    Friusername: f.Friusername,
  })
  if ok1 && ok2 {
    _, err1 := o.Delete(&Friendlist{Fid: fid1})
    _, err2 := o.Delete(&Friendlist{Fid: fid2})
    if err1 == nil && err2 == nil{
      return true
    }
  }
  return false
}


/************************************************
 * invitation table
 ***********************************************/
type ivttTbl struct {
  name string
}

func (ivtt *ivttTbl) Insert(i *JsIvtt) bool {
  o := orm.NewOrm()
  o.Using("default")

  invite := Invitation{
    Sender: i.Sender,
    Receiver: i.Receiver,
    Msg: i.Msg,
  }

  if ivtt.Exist(i.Sender, i.Receiver) == false && i.Sender != i.Receiver{
    _, err := o.Insert(&invite)
    if err == nil {
      return true
    }
  }

  return false
}

//判断邀请函是否存在
func (ivtt *ivttTbl) Exist(sdr string, rec string) bool {
  o := orm.NewOrm()
  o.Using("defalut")

  var invites []Invitation
  num, err := o.QueryTable("invitation").Filter("sender", sdr).Filter("receiver", rec).All(&invites)
  if err == nil {
    if num == 1 {
      return true
    }
  }
  return false
}

/*
 * name: username
 * key: "sender" or "receiver"
 */
func (ivtt *ivttTbl) ReadList(result *[]JsIvtt, name string, key string) bool {
  o := orm.NewOrm()
  o.Using("default")

  var invites []Invitation

  _, err := o.QueryTable("invitation").Filter(key, name).All(&invites)
  if err == nil {
    for _, v := range invites {
      var item JsIvtt
      item.Sender = v.Sender
      item.Receiver = v.Receiver
      item.Msg = v.Msg
      *result = append(*result, item)
    }
    return true
  }

  return false
}

func (ivtt *ivttTbl) ReadId(sdr string, rec string) (int, bool) {
  o := orm.NewOrm()
  o.Using("default")

  var invite []Invitation

  num, err := o.QueryTable("invitation").Filter("sender", sdr).Filter("receiver",  rec).All(&invite)
  if err == nil {
    if num == 1 {
      return invite[0].Iid, true
    }
  }

  return -1, false

}

//单项删除
func (ivtt *ivttTbl) Delete(sdr string, rec string) bool {
  o := orm.NewOrm()
  o.Using("default")

  id, ok := ivtt.ReadId(sdr, rec)

  if ok {
    _, err := o.Delete(&Invitation{Iid: id})
    if err == nil {
      return true
    }
  }
  return false
}

/***************************************************
 * user table
 **************************************************/
type userTbl struct{
  name string
}

func (user *userTbl) Exist(username string) bool {
  o := orm.NewOrm()
  o.Using("default")

  usr := User{Username: username}
  err := o.Read(&usr, "Username")

  if err == orm.ErrNoRows || err == orm.ErrMissPK {
    return false
  } else {
    return true
  }

}

func (user *userTbl) Insert(u *JsUser) bool {
  if user.Exist(u.Username) {
    return false
  }

  o := orm.NewOrm()
  o.Using("default")
  var usr User
  usr.Username = u.Username
  pwd, err := scrypt.Key([]byte(u.Password), Salt, 16384, 8, 1, 32)
  if err != nil {
    return false
  }
  usr.Password = string(pwd)
  usr.Isactive = 0

  _, err = o.Insert(&usr)
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

func (user *userTbl) Verify(u *JsUser) bool {
  o := orm.NewOrm()
  o.Using("default")

  usr := User{Username: u.Username}
  err := o.Read(&usr, "Username")

  if err == orm.ErrNoRows || err == orm.ErrMissPK {
    return false
  } else {
    if usr.Isactive == 1 {
      return false
    }

    pwd, err := scrypt.Key([]byte(u.Password), Salt, 16384, 8, 1, 32)
    if err != nil {
      return false
    }

    return byteSliceEqual(pwd, []byte(usr.Password))
  }
}

func (user *userTbl) SetActive(u string) bool {
  o := orm.NewOrm()
  o.Using("default")

  usr := User{Username: u}
  err := o.Read(&usr, "Username")

  if err == orm.ErrNoRows || err == orm.ErrMissPK {
    return false
  } else {
    usr.Isactive = 1
    if _, err = o.Update(&usr); err == nil {
      return true
    } else {
      return false
    }
  }

}

func (user *userTbl) SetUnActive(u string) bool {
  o := orm.NewOrm()
  o.Using("default")

  usr := User{Username: u}
  err := o.Read(&usr, "Username")

  if err == orm.ErrNoRows || err == orm.ErrMissPK {
    return false
  } else {
    usr.Isactive = 0
    if _, err = o.Update(&usr); err == nil {
      return true
    } else {
      return false
    }
  }
}

/***************************************************
 * userinfo table
 **************************************************/
type userinfoTbl struct {
  name string
}

func (uif *userinfoTbl) Insert(u *JsUif) bool {
  if uif.Exist(u.Username, "Username") {
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

func (uif *userinfoTbl) Read(u *JsUif, key string) bool {
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

func (uif *userinfoTbl) Update(u *JsUif) bool {
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

func (uif *userinfoTbl) UpdateIcon(u *JsUif) bool {
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

func (uif *userinfoTbl) Exist(content string, key string) bool {
  o := orm.NewOrm()
  o.Using("default")

  var usr Userinfo
  switch key {
  case "Username":
    usr.Username = content
  case "Nickname":
    usr.Nickname = content
  case "Email":
    usr.Email = content
  }

  err := o.Read(&usr, key)

  if err == orm.ErrNoRows || err == orm.ErrMissPK {
    return false
  } else {
    return true
  }
}

func (uif *userinfoTbl) ReadFifList(u *[]JsUif, f *[]JsFl) bool {
  o := orm.NewOrm()
  o.Using("default")

  for _, v := range *f {
    item := JsUif{Username: v.Friusername}
    ok := uif.Read(&item, "Username")
    if !ok { return false }
    *u = append(*u, item)
  }

  return true
}
