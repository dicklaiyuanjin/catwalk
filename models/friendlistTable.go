package models

import (
  "github.com/astaxie/beego/orm"
)

func IsFriListExist(f *JsFl) bool {
  o := orm.NewOrm()
  o.Using("default")

  var fri Friendlist
  num, err := o.QueryTable("friendlist").Filter("username", f.Username).Filter("friusername", f.Friusername).All(&fri)
  if err == nil {
    if num == 1 {
      return true
    }
  }
  return false
}

func InsertFriendList(f *JsFl) bool{
  o := orm.NewOrm()
  o.Using("default")

  fri := Friendlist{
    Username: f.Username,
    Friusername: f.Friusername,
  }

  if IsFriListExist(f) == false && f.Username != f.Friusername {
    _, err := o.Insert(&fri)
    if err == nil {
      return true
    }
  }

  return false
}

func ReadFriendList(f *[]JsFl, username string) bool {
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

func ReadFriListId(f *JsFl) (int, bool) {
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

func DeleteFriendList(f *JsFl) bool {
  o := orm.NewOrm()
  o.Using("default")

  fid, ok := ReadFriListId(f)

  if ok {
    if _, err := o.Delete(&Friendlist{Fid: fid}); err == nil {
      return true
    }
  }
  return false

}